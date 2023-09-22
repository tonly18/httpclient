package main

import (
	"fmt"
	"httpclient/httpclient"
	"runtime"
)

func main() {
	for i := 0; i < 10000; i++ {
		httpClient := httpclient.NewHttpClient(&httpclient.Config{})
		resp, err := httpClient.Get("http://192.168.1.45:6000/v1/test", nil).Do()
		fmt.Println("err1:::::", err)

		data, err := resp.GetData()
		fmt.Println("err2:::::", err)
		fmt.Println("resp:::::", string(data))
		fmt.Println("nasdfasdasdf:::::", runtime.NumGoroutine())
		fmt.Println("")
	}

	fmt.Println("main")
}
