# Very Simple Http Client For Requests
Integrated with Heimdall and FastHTTP

## Installation

```bash
go get github.com/hackincloud-com/hic-client
```

## Usage

Just like this

```go
package main

import (
	httpclient "github.com/hackincloud-com/hic-client"
)
func main() {
	// if POST
	client := &httpclient.ReqStruct{
		Url:     "https://google.com",
		Method:  "POST",
		Body:    io.ReaderBody,
		Timeout: 60,
	}

	res := client.DoRequests()
	if res.Error != nil {
		cobra.CheckErr(res.Error)
	}
    // if GET
    client := &httpclient.ReqStruct{
		Url:     "https://google.com",
		Method:  "GET",
		Timeout: 60,
	}

	res := client.DoRequests()
	if res.Error != nil {
		cobra.CheckErr(res.Error)
	}
}
```
