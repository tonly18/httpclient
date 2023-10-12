package httpclient_test

import (
	"fmt"
	"github.com/tonly18/httpclient"
	"testing"
	"time"
)

func TestHttpClient(t *testing.T) {
	httpClient := httpclient.NewHttpClient(&httpclient.Config{})
	for i := 0; i < 100; i++ {
		go func() {
			resp, err := httpClient.NewRequest("POST", "http://192.168.1.30:6000/v1/test", nil).Do()
			fmt.Println("err:::::", err)
			fmt.Println("resp.header:::::", resp.GetHeaderCode())

			data, err := resp.GetData()
			fmt.Println("err:::::", err)
			fmt.Println("resp:::::", string(data))
			fmt.Println("")
		}()
	}

	time.Sleep(6 * time.Second)

	fmt.Println("main")
}
