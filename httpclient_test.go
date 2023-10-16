package httpclient_test

import (
	"fmt"
	"github.com/tonly18/httpclient"
	"math/rand"
	"testing"
	"time"
)

func TestHttpClient(t *testing.T) {

	for i := 0; i < 1000; i++ {
		go func() {
			httpClient := httpclient.NewHttpClient(&httpclient.Config{})
			resp, err := httpClient.NewRequest("POST", "http://192.168.1.30:6000/v1/test", nil).SetHeader(map[string]any{
				"name": "Sam",
				"age":  rand.New(rand.NewSource(time.Now().UnixNano() + int64(rand.Int()))).Int(),
			}).Do()
			//fmt.Println("err:::::", err)
			//fmt.Println("resp.header:::::", resp.GetHeaderCode())

			data, err := resp.GetData()
			fmt.Println("err:::::", err)
			fmt.Println("resp:::::", string(data))
			fmt.Println("")
			time.Sleep(time.Nanosecond * 1)
		}()
	}

	time.Sleep(10 * time.Second)

	fmt.Println("main")
}
