# cs-reboot-info

This is a Rackspace tool to identify Cloud Servers that have a scheduled automated reboot window. Cloud Servers may have a scheduled reboot in the case of routine or critical system maintenance.

The tool source is OS independent (written in Go) and binaries are available for Windows, Mac OS X, and Linux.

# How it works

*cs-reboot-info* queries the Rackspace Cloud US and UK cloud infrastructures (both First and Next Generation). It identifies any Cloud Servers with a metadata key named *"rax:reboot_window"*. This key carries a value that shows the start and end times of the scheduled reboot window for the Cloud Server.

The format of the metadata key is:

```
| Key               | Value  (example)                          |
|-------------------|-------------------------------------------|
| rax:reboot_window | 2014-01-28T00:00:00Z;2014-01-28T03:00:00Z |
```

The value is a semi-colon separated time range, in UTC format.

The tool outputs a list of Cloud Servers that have scheduled reboot windows in a tabular format. Results can optionally be saved to a CSV file.

**Note:** Only Cloud Servers with a scheduled reboot window will be listed. If a Cloud Server is not listed, no automated reboots are scheduled for it.


## Installation - Binaries

| Plaform        | Download links |
| -------------- | -------------- |
| Windows 386    | [`cs-reboot-info_windows_386.exe`](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_windows_386.exe) ([SHA1](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_windows_386.exe.sha1)) |
| Windows AMD64  | [`cs-reboot-info_windows_amd64.exe`](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_windows_amd64.exe) ([SHA1](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_windows_amd64.exe.sha1)) |
| Mac OS X 386   | [`cs-reboot-info_darwin_386`](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_darwin_386) ([SHA1](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_darwin_386.sha1)) |
| Mac OS X AMD64 | [`cs-reboot-info_darwin_amd64`](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_darwin_amd64) ([SHA1](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_darwin_amd64.sha1)) |
| FreeBSD 386    | [`cs-reboot-info_freebsd_386`](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_freebsd_386) ([SHA1](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_freebsd_386.sha1)) |
| FreeBSD AMD64  | [`cs-reboot-info_freebsd_amd64`](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_freebsd_amd64) ([SHA1](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_freebsd_amd64.sha1)) |
| FreeBSD ARM    | [`cs-reboot-info_freebsd_arm`](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_freebsd_arm) ([SHA1](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_freebsd_arm.sha1)) |
| Linux 386      | [`cs-reboot-info_linux_amd64`](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_linux_386) ([SHA1](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_linux_386.sha1)) |
| Linux AMD64    | [`cs-reboot-info_linux_amd64`](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_linux_amd64) ([SHA1](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_linux_amd64.sha1)) |
| Linux ARM      | [`cs-reboot-info_linux_arm`](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_linux_arm) ([SHA1](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_linux_arm.sha1)) |
| NetBSD 386     | [`cs-reboot-info_netbsd_386`](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_netbsd_386) ([SHA1](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_netbsd_386.sha1)) |
| NetBSD AMD64   | [`cs-reboot-info_netbsd_amd64`](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_netbsd_amd64) ([SHA1](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_netbsd_amd64.sha1)) |
| NetBSD ARM     | [`cs-reboot-info_netbsd_arm`](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_netbsd_arm) ([SHA1](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_netbsd_arm.sha1)) |
| OpenBSD 386    | [`cs-reboot-info_openbsd_386`](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_openbsd_386) ([SHA1](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_openbsd_386.sha1)) |
| OpenBSD AMD64  | [`cs-reboot-info_openbsd_amd64`](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_openbsd_amd64) ([SHA1](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_openbsd_amd64.sha1)) |
| Plan 9 386     | [`cs-reboot-info_plan9_386`](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_plan9_386) ([SHA1](https://a4fae0f0d6cf4cc92acd-d6ce857812540f8fb39144d83ca6538f.ssl.cf5.rackcdn.com/08c510b/cs-reboot-info_plan9_386.sha1)) |

After you download the binary, place it anywhere on your `${PATH}` (or `%PATH%`) and rename it to `cs-reboot-info`. On Mac or Linux, you'll also need to `chmod +x cs-reboot-info` to make it executable.

## Using the tool

**Usage:**
```bash
cs-reboot-info [--csv] username apikey
```
*Username* and *apikey* are required arguments, and are the same credentials you normally use with the Rackspace Cloud API.

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
generation, region, server_uuid, server_name, reboot_window_start_UTC, reboot_window_end_UTC, reboot_window_start_local, reboot_window_end_local
First Gen,DFW,d7b47a17-1552-4dcd-8b7b-831fddd73c42,MyFGServer,01 Jan 00:00,01 Jan 00:00,01 Jan 00:00,01 Jan 00:00
Next Gen,IAD,4da4a108-99c3-448a-8791-0e3fa81cbc98,MyNGServer,01 Jan 00:00,01 Jan 00:00,01 Jan 00:00,01 Jan 00:00
```

## Building from source

### Prerequisites

* A working [Go installation](https://golang.org/doc/install).
* A healthy [Go workspace](https://golang.org/doc/code.html#Organization).

Clone the github repo as normal, and run:

```bash
go get github.com/rackerlabs/cs-reboot-info
cd ${GOPATH}/src/github.com/rackerlabs/cs-reboot-info

# Fetch dependencies into your ${GOPATH}
go get .

# Build and install the binary to ${GOPATH}/bin
go install .
```
