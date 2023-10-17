package httpclient_test

import (
	"fmt"
	"github.com/tonly18/httpclient"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestHttpClient(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			httpClient := httpclient.NewHttpClient(&httpclient.Config{})
			resp, err := httpClient.Get("http://localhost:8080/n/v1/test", map[string]any{
				"fruit": "mango-123",
			}).SetHeader(map[string]any{
				"proxy_id":  100,
				"server_id": 2,
				"client_ip": "192.168.1.48",
				"user_id":   21,
				"trace_id":  15821793512,
			}).Do()
			fmt.Println("err:::::", err)
			fmt.Println("resp.header:::::", resp.GetHeaderCode())

			data, err := resp.GetData()
			fmt.Println("err:::::", err)
			fmt.Println("resp:::::", string(data))
			fmt.Println("")
			time.Sleep(time.Nanosecond * 1)
		}()
	}

	time.Sleep(6 * time.Second)

	fmt.Println("main")
}

func TestClientGet(t *testing.T) {
	rawurl := "http://localhost:8080/n/v1/test?fruit=mango"
	httpClient := httpclient.NewHttpClient(&httpclient.Config{})
	resp, err := httpClient.NewRequest(http.MethodGet, rawurl, nil).SetHeader(map[string]any{
		"proxy_id":  100,
		"server_id": 2,
		"client_ip": "192.168.1.48",
		"user_id":   21,
		"trace_id":  15821793512,
	}).Do()
	data, err := resp.GetData()
	fmt.Println("err:::::", err)
	fmt.Println("resp:::::", string(data))

	fmt.Println("main")
}

func TestClientPost(t *testing.T) {
	// "content-type": "application/x-www-form-urlencoded",
	body := "fruit=orange&color=red"
	bodyByte := []byte(body)

	//"content-type": "application/json"
	//body := map[string]any{
	//	"fruit":   "orange-22",
	//	"color":   "color-22",
	//	"user_id": 22,
	//}
	//bodyByte, _ := json.Marshal(body)

	rawurl := "http://localhost:8080/n/v1/test/post?sex=3"
	httpClient := httpclient.NewHttpClient(&httpclient.Config{})
	resp, err := httpClient.NewRequest(http.MethodPost, rawurl, bodyByte).SetHeader(map[string]any{
		//"content-type": "application/x-www-form-urlencoded",
		//"content-type": "application/json",
		//"content-type": "application/octet-stream",
		"proxy_id":  110,
		"server_id": 22,
		"client_ip": "192.168.1.48",
		"user_id":   21,
		"trace_id":  15821793512,
	}).Do()
	if err != nil {
		log.Fatal("err:::::::", err)
	}
	data, err := resp.GetData()
	fmt.Println("err:::::", err)
	fmt.Println("resp:::::", string(data))

	fmt.Println("main")
}
