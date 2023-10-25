# httpclient
httpclient是模拟http请求包，支持链接复用，使用缓存池性能更好。



## Installation

```bash
go get -u github.com/tonly18/httpclient
```

## Getting Started

### Simple httpclient Example

For simple httpclient, import the global httpclient package **github.com/tonly18/httpclient**



```go
package main

import (
	"fmt"
	"github.com/tonly18/httpclient"
	"log"
)

func main() {
	httpClient := httpclient.NewClient(&httpclient.Config{})
	for i := 0; i < 10; i++ {
		resp, err := httpClient.Get("http://www.baidu.com", nil).SetHeader(map[string]any{
			"name": "Sam",
			"age":  18,
		}).Do()

		if err != nil {
			log.Fatalf(`%v`, err)
		}
		
		data, _ := resp.GetData()
		fmt.Println("resp:", string(data))

	}
}
```
