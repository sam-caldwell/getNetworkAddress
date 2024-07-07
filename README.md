getNetworkAddress
=================

## Description

Create package and command-line tool to get the network address from a given CIDR address.

## Usage

Run this command:
```shell
go run cmd/main.go "10.1.2.3/24"
```
which will produce the network address:
```shell
10.1.2.0
```
If the `-trim` flag is added to the call, the result will be--
```shell
10.1.2
```