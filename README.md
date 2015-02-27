#cs-reboot-info

This is a Rackspace tool to identify Cloud Servers that have a scheduled automated reboot window. Cloud Servers may have a scheduled reboot in the case of routine or critical system maintenance. 

The tool source is OS independent (written in Go) and binaries are available for Windows, Mac OS X, and Linux. 

# How it works

*cs-reboot-info* queries the Rackspace Cloud US and UK cloud infrastructures (both First and Next Generation). It identifies any Cloud Servers with a metadata key named *"rax:reboot_window"*. This key carries a value that shows the start and end times of the scheduled reboot window for the Cloud Server. 

The format of the metadata key is:


| Key | Value  (example)|
|--------------------------|
| rax:reboot_window | 2014-01-28T00:00:00Z;2014-01-28T03:00:00Z |


The value is a semi-colon separated time range, in UTC format. 

The tool outputs a list of Cloud Servers that have scheduled reboot windows in a tabular format. Results can optionally be saved to a CSV file. 

**Note:** Only Cloud Servers with a scheduled reboot window will be listed. If a Cloud Server is not listed, no automated reboots are scheduled for it. 


## Installation - Binaries

The *cs-reboot-info* tool is a single binary package (per platform), with no additional dependencies.  Just download and run!

### Linux, Mac OS X 

If you are a Linux or OS X user, you can download the binary directly from Github using wget:

**Linux**: 
```bash
wget https://github.com/cs-reboot-info/bin/linux/cs-reboot-info
```
**Mac OS X**: 
```bash
wget https://github.com/cs-reboot-info/bin/osx/cs-reboot-info
```

### Windows 

If you are a Windows user, you can download the tool from: https://github.com/cs-reboot-info/bin/windows/cs-reboot-info

## Building from source

### Prerequisites
* A working Go installation: https://golang.org/doc/install
* GopherCloud: https://github.com/rackspace/gophercloud

Clone the github repo as normal, and run:

```bash
go cs-reboot-info.go
```

## Using the tool

**Usage:**
```bash
cs-reboot-info [--csv] username apikey
```
*Username* and *apikey* are required arguments, and are the same credentials you normally use with the Rackspace Cloud API (or to log in the Cloud Control Panel). 

*--csv*: Optional, used to specify that you also want the results stored in a CSV file titled **cs-reboot-info.csv** in the same directory as the tool. 


### Sample output: Table (default)

```
| Type            | Server ID                            | Server Name          | Reboot Window (UTC)         | Reboot Window (Local)       |
| --------------- | ------------------------------------ | -------------------- | --------------------------- | --------------------------- |
| First Gen (IAD) | 8c65cb68-0681-4c30-bc88-6b83a8a26aee | Gophercloud-pxpGGuey | 01 Jan 00:00 - 01 Jan 00:00 | 01 Jan 01:00 - 01 Jan 01:00 |
| First Gen (IAD) | 8c65cb68-0681-4c30-bc88-6b83a8a26aee | Gophercloud-pxpGGuey | 01 Jan 00:00 - 01 Jan 00:00 | 01 Jan 01:00 - 01 Jan 01:00 |
```

### Sample output: CSV
```
generation, region, server uuid, server name, reboot time window UTC, reboot time window local
first,DFW,d7b47a17-1552-4dcd-8b7b-831fddd73c42,MyFGServer,01 Jan 00:00-01 Jan 00:00,01 Jan 00:00-01 Jan 00:00
next,IAD,4da4a108-99c3-448a-8791-0e3fa81cbc98,MyNGServer,01 Jan 00:00-01 Jan 00:00,01 Jan 00:00-01 Jan 00:00
```



