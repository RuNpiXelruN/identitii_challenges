# Identitii Rate Limit Challenge

## Build Setup

``` bash
# install dependencies
Fetch the godep package by running,

go get -u github.com/tools/godep

followed by,

godep restore

# To run
go run main.go, to use default job & worker values

go run main.go --help, to view what cli flags the program takes

# To build binary
go install, followed by
go build

# To run binary
./rate_limit
```