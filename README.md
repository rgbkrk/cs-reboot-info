#cs-reboot-info

This is a Rackspace tool to identify Cloud Servers that have a scheduled automated reboot window. Cloud Servers may have a scheduled reboot in the case of routine or critical system maintenance. 

The tool source is OS independent (written in Go) and binaries are available for Windows, MacOS, and Linux. 

# How it works

*cs-reboot-info* queries the Rackspace Cloud US and UK cloud infrastructures (both First and Next Generation). It identifies any Cloud Servers with a metadata key named *"rax:reboot_window"*. This key carries a value that shows the start and end times of the scheduled reboot window for the Cloud Server. 

The format of the metadata:

| Key | Value  (example)|
|--------------------------|
| rax:reboot_window | 2014-01-28T00:00:00Z;2014-01-28T03:00:00Z |
The value is a semi-colon separated time range, in UTC format. 

The tool outputs a list of Cloud Servers that have scheduled reboot windows in a tabular format. 

**Note:** Only Cloud Servers with a scheduled reboot window will be listed. If a Cloud Server is not listed, no automated reboots are scheduled for it. 

Results can optionally be saved to a CSV file called *cs-reboot-info-output.csv*, in the following format:
    
*generation, region, server uuid, server name, reboot time window UTC, reboot time window local*


## Installation - Binaries

The *cs-reboot-info* tool is a single binary package (per platform), with no additional dependencies.  

### Linux, Mac OSX 

If you are a Linux or OSX user, you can download the binary directly from Github using wget:

**Linux**: 
```bash
wget https://github.com/cs-reboot-info/bin/linux/cs-reboot-info
```
**Mac OS X**: 
```bash
wget https://github.com/cs-reboot-info/bin/osx/cs-reboot-info
```

### Windows 

If you are a Windows user, you can download from the following link in your browser: https://github.com/cs-reboot-info/bin/windows/cs-reboot-info

## Building from source

### Prerequisites
* A working Go installation: https://golang.org/doc/install
* GopherCloud: https://github.com/rackspace/gophercloud

Clone the github repo as normal, and run:

```bash
go cs-reboot-info.go
```

