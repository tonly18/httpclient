package main

import (
	"fmt"
	"httpclient/httpclient"
)

func main() {
	httpClient := httpclient.NewHttpClient(&httpclient.Config{})
	resp, err := httpClient.Get("http://192.168.1.30:6000/v1/test", nil).Do()
	fmt.Println("err1:::::", err)

	data, err := resp.GetData()
	fmt.Println("err2:::::", err)
	fmt.Println("resp:::::", len(data))
	fmt.Println("resp:::::", string(data))

	fmt.Println("main")
}
