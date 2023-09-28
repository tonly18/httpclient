package httpclient_test

import (
	"fmt"
	"github.com/tonly18/httpclient"
	"testing"
	"time"
)

func TestNewHttpClient(t *testing.T) {
	httpClient := httpclient.NewHttpClient(&httpclient.Config{})
	for i := 0; i < 10; i++ {
		go func() {
			resp, err := httpClient.NewRequest("POST", "http://192.168.1.45:6000/v1/test", nil).Do()
			fmt.Println("err:::::", err)

			data, err := resp.GetData()
			fmt.Println("err:::::", err)
			fmt.Println("resp:::::", string(data))
			fmt.Println("")
		}()
	}

	time.Sleep(2 * time.Second)

	fmt.Println("main")
}
