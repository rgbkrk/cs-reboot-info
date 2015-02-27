package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jrperritt/gophercloud/rackspace"
	rsV1Servers "github.com/jrperritt/gophercloud/rackspace/compute/v1/servers"
	"github.com/rackspace/gophercloud"
	osV2Servers "github.com/rackspace/gophercloud/openstack/compute/v2/servers"
	"github.com/rackspace/gophercloud/pagination"
	rsV2Servers "github.com/rackspace/gophercloud/rackspace/compute/v2/servers"
	"github.com/rackspace/gophercloud/rackspace/identity/v2/tokens"
)

const metadataTimeFmt = "2006-01-02T15:04:05Z"

func main() {
	outputToCSV := *flag.Bool("csv", false, "Output a CSV file")
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Println("You must supply a username and API key after all other args.")
		os.Exit(1)
	}

	username, apiKey := flag.Arg(0), flag.Arg(1)

	opts := gophercloud.AuthOptions{
		Username: username,
		APIKey:   apiKey,
	}

	provider, err := rackspace.AuthenticatedClient(opts)
	if err != nil {
		fmt.Printf("Unable to authenticate: %v", err)
	}

	regions, fg := Regions(provider, opts)

	fmt.Printf("Regions with a compute endpoint: %s\n", strings.Join(regions, ", "))
	if fg {
		fmt.Println("You do have a first-gen endpoint, too.")
	}

	var entries []entry

	// Iterate through regions with an NG compute endpoint. Collect data about each server.
	for _, region := range regions {
		compute, err := rackspace.NewComputeV2(provider, gophercloud.EndpointOpts{
			Region: region,
		})
		if err != nil {
			fmt.Printf("Unable to locate a v2 compute endpoint in region %s: %v\n", region, err)
			continue
		}

		err = rsV2Servers.List(compute, nil).EachPage(func(page pagination.Page) (bool, error) {
			s, err := osV2Servers.ExtractServers(page)
			if err != nil {
				return false, err
			}

			for _, server := range s {
				md, err := osV2Servers.Metadatum(compute, server.ID, "rax:reboot_window").Extract()
				if err != nil {
					fmt.Printf("Unable to retrieve rax:reboot_window metadatum for server %s: %v\n", server.ID, err)
					continue
				}

				windowString, ok := md["rax:reboot_window"]
				if !ok {
					fmt.Printf("Metadatum rax:reboot_window was not present in the result for server %s.\n", server.ID)
					continue
				}

				// Expected format: 2014-01-28T00:00:00Z;2014-01-28T03:00:00Z

				parts := strings.Split(windowString, ";")
				if len(parts) != 2 {
					fmt.Printf("Unexpected metadatum format for server %s: %s\n", server.ID, windowString)
					continue
				}

				start, err := time.Parse(metadataTimeFmt, parts[0])
				if err != nil {
					fmt.Printf("Unable to parse window start time for server %s: %s\n", server.ID, parts[0])
					continue
				}

				end, err := time.Parse(metadataTimeFmt, parts[1])
				if err != nil {
					fmt.Printf("Unable to parse window end time for server %s: %s\n", server.ID, parts[1])
					continue
				}

				entry := entry{
					Server:      server,
					Region:      region,
					GenType:     "Next Gen",
					WindowStart: start,
					WindowEnd:   end,
				}
				entries = append(entries, entry)
			}

			return true, nil
		})
	}

	// Iterate through regions with an FG compute endpoint. Collect data about each server.
	compute, err := rackspace.NewComputeV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Printf("Unable to locate a v1 compute endpoint: %v\n", err)
	}

	err = rsV1Servers.List(compute, rsV1Servers.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		s, err := osV2Servers.ExtractServers(page)
		if err != nil {
			return false, err
		}

		for _, server := range s {
			windowString, ok := server.Metadata["rax:reboot_window"]
			if !ok {
				fmt.Printf("Metadatum rax:reboot_window was not present in the result for server %s.\n", server.ID)
				continue
			}

			// Expected format: 2014-01-28T00:00:00Z;2014-01-28T03:00:00Z

			parts := strings.Split(windowString.(string), ";")
			if len(parts) != 2 {
				fmt.Printf("Unexpected metadatum format for server %s: %s\n", server.ID, windowString)
				continue
			}

			start, err := time.Parse(metadataTimeFmt, parts[0])
			if err != nil {
				fmt.Printf("Unable to parse window start time for server %s: %s\n", server.ID, parts[0])
				continue
			}

			end, err := time.Parse(metadataTimeFmt, parts[1])
			if err != nil {
				fmt.Printf("Unable to parse window end time for server %s: %s\n", server.ID, parts[1])
				continue
			}

			entry := entry{
				Server:      server,
				Region:      "DFW",
				GenType:     "First Gen",
				WindowStart: start,
				WindowEnd:   end,
			}
			entries = append(entries, entry)
		}

		return true, nil
	})

	// Pull the metadata key

	if outputToCSV {
		outputCSV(entries)
	} else {
		outputTabular(entries)
	}
}

// Regions acquires the service catalog and returns a slice of every region that contains a next-gen
// server endpoint, and a boolean indicating whether or not this customer has access to FG servers.
func Regions(provider *gophercloud.ProviderClient, opts gophercloud.AuthOptions) ([]string, bool) {
	service := rackspace.NewIdentityV2(provider)

	result := tokens.Create(service, tokens.WrapOptions(opts))
	catalog, err := result.ExtractServiceCatalog()
	if err != nil {
		fmt.Printf("Unable to retrieve the service catalog: %v\n", err)
		os.Exit(1)
	}

	var regions []string
	var fg bool
	for _, entry := range catalog.Entries {
		if entry.Type == "compute" {
			for _, endpoint := range entry.Endpoints {
				if endpoint.Region == "" {
					fg = true
				} else {
					regions = append(regions, endpoint.Region)
				}
			}
		}
	}
	return regions, fg
}
