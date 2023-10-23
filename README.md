# httpclient
http client是模拟http请求的包


go get -u github.com/tonly18/httpclient


### 注：<br>
#### 一个httpclient对应一个request，如果需要多次循环调用，则需要每次新建httpclient。
#### 例如:

```go
package main

import (
	"github.com/tonly18/httpclient"
)

func main() {
	for i:=0; i<10; i++{
		httpClient := httpclient.NewHttpClient(&httpclient.Config{})
		httpClient.Get("http://www.baidu.com", nil).Do()
		
	}
}
```
