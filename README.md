## Aliyun SDK for go

## How to

``` sh
git clone ...

# initial
git submodule update --init
gradle 

# update
git submodule update --remote
gradle
```

## Issues

* Just PRC 

## Usage

```go
package main

import (
    "os"
    "fmt"
    "time"

    "github.com/morlay/goaliyun"
    "github.com/morlay/goaliyun/clients/vpc"
)

func main() {
    client := goaliyun.NewClientWithCredential(
      	os.Getenv("ALIYUN_ACCESS_KEY"),
      	os.Getenv("ALIYUN_SECRET_KEY"),

        // optional config timeout
        goaliyun.ClientOptionWithTimeout(1 * time.Hour),
        // log each request        
        goaliyun.ClientOptionWithTransports(
        	goaliyun.NewLogTransport(),
        	// or define your own transport
        ),
    )
    
    req := vpc.DescribeRegionsRequest{
        RegionId: "cn-beijing",
    }
    
    resp, err := req.Invoke(client)
    
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("%v", resp.Regions)
}
```
