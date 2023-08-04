package main

import (
	"fmt"
	"httpclient/httpclient"
)

func main() {
	httpClient := httpclient.NewHttpClient(&httpclient.Config{})
	resp, err := httpClient.Get("https://baidu.com", nil).Do()
	fmt.Println("err:::::", err)

	data, _ := resp.GetData()
	fmt.Println("resp::::", string(data))

	fmt.Println("main")
}
