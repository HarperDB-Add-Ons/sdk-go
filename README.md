# HarperDB SDK for Go

[![PkgGoDev](https://pkg.go.dev/badge/github.com/HarperDB-Add-Ons/sdk-go)](https://pkg.go.dev/github.com/HarperDB-Add-Ons/sdk-go)

This is the Go SDK for HarperDB.

## Requirements

- >= Go 1.18

## Installation

```
go get github.com/HarperDB-Add-Ons/sdk-go
```

## Quickstart

```go
client := harperdb.NewClient("http://localhost:9925", "HDB_ADMIN", "password")
client.CreateSchema("dog")
```