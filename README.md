# Golang-based Project for Aliyun Domain Dynamic DNS Integration

## Environmental Requirements

The environment requires `golang 1.21.*`, with the recommended version being `golang 1.21.9`.

## Setting up Domestic Golang Proxy

```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

## Installing Project Dependencies

```bash
go mod tidy
```

## Compiling Code

Note that code compiled on Windows can only be executed on Windows. To run on Linux, recompilation is required on a Linux system. The compilation on Windows generates an .exe executable file. Use the following script to compile (on Windows, install git and execute with git bash):

```bash
./build.sh
```

Upon completion, an executable file will appear in the project's `bin` directory with different filenames for each system:

- Windows: greateme_ddns.exe
- Linux/MacOS: greateme_ddns
  Additionally, a `config.ini` configuration file will be generated in the `bin/conf` directory. This configuration file needs to be edited as follows:
- `accessKeyId`: Change to your Aliyun `accessKey`
- `accessKeySecret`: Change to your Aliyun `accessKeySecret`
- `domainEndpoint`: The domain query Endpoint, defaults to Hangzhou and does not require modification
- `dnsEndpoint`: The DNS Endpoint, defaults to Shenzhen and can be modified based on annotations in the configuration file and geographical location
- `domainList`: A list of domains, separated by commas
- `dnsType`: The resolution type, can only be ipv4 or ipv6, defaults to ipv4 (note: must be lowercase and cannot be uppercase)
- `type`: Execution type, options are: single and repetition. single: Execute only once, requires coordination with the system's cron job. repetition: Execute repeatedly, requires coordination with the durationMinute configuration item
- `durationMinute`: How often to update (in minutes), defaults to ten minutes and can be left unchanged

## Running the Code

You can directly execute the executable file in the bin directory.
