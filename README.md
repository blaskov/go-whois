# go-whois

## About

*go-whois* is a simple Go module for interaction with [WHOIS] servers.

### Features

* Support for [RFC 3912][WHOIS]
* Support for \*.whois-servers.net CNAMEs to determine the correct WHOIS
  server to be queried
* Allows the developer to specify WHOIS server address and port, if needed
* Works most of the time ;)

### Disclaimer

This library is still in early stage of development, and is my first attempt
to write a Go library, so it may not be ready for production use yet. On the
other hand, pull requests and constructive criticism are welcome.

## Usage

### Requirements

Before using this module you should install it into your environment with the
command `go get github.com/blaskov/go-whois`.

### Functions

Currently only one function is exported by the *go-whois* module:

```go
func Whois(query string, params map[string]string) (result string, err error)
```

#### Arguments

1. `query` -- query string sent to the remote server, a domain name or IP
    address in most cases
2. `params` -- a map that currently accepts two keys:
  * `host` -- hostname or IP address of the WHOIS server to be contacted
  * `port` -- TCP port of the WHOIS server to be contacted

#### Return values

1. `result` -- a string returned by the remote server in reply of the sent
   query
2. `err` -- in case of failure, this variable is populated with an error
   message

### Example

WHOIS client in less than 50 lines of code:

```go
package main

import (
	"flag"
	"fmt"
	"github.com/blaskov/go-whois"
    "os"
)

var host = flag.String("host", "", "Send query directly to the specified server")
var port = flag.String("port", "", "Specify a port number to use when querying a WHOIS server")

func main() {
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: %s [arguments] <domain>\n\nAvailable arguments:\n", os.Args[0])
        flag.PrintDefaults()

        os.Exit(1)
    }

	flag.Parse()
	query := flag.Arg(0)
    params := make(map[string]string)

    if query == "" {
        flag.Usage()
    }

    if *host != "" {
        params["host"] = *host
    }

    if *port != "" {
        params["port"] = *port
    }

	info, err := whois.Whois(query, params)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(info)
	}
}
```

[WHOIS]: http://tools.ietf.org/html/rfc3912
