package httpclient_test

import (
	"fmt"
	"github.com/tonly18/httpclient"
	"log"
	"net/http"
	"runtime"
	"testing"
	"time"
)

func TestHttpClient(t *testing.T) {
	fmt.Println("goroutineNum:", runtime.NumGoroutine())

	httpClient := httpclient.NewClient(&httpclient.Config{})
	for i := 0; i < 1000; i++ {
		go func(m int) {
			for x := 0; x < 100; x++ {
				resp, err := httpClient.Get("http://192.168.1.30:8080/n/v1/test", map[string]any{
					"fruit": "mango-123",
				}).SetHeader(map[string]any{
					"proxy_id":  100,
					"server_id": 2,
					"client_ip": "192.168.1.48",
					"user_id":   21,
					"trace_id":  15821793512,
				}).Do()
				if err != nil {
					fmt.Println("err:::::", err)
				} else {
					fmt.Println("resp.header:::::", resp.GetHeaderCode())
					data, err := resp.GetData()
					if m%10 == 0 {
						fmt.Println("err:::::", err)
						fmt.Println("resp:::::", string(data))

						fmt.Println("goroutineNum:", runtime.NumGoroutine())
						fmt.Println("")
					}
				}

				time.Sleep(time.Nanosecond * 1000)
			}
		}(i)
	}

	fmt.Println("goroutineNum-2:", runtime.NumGoroutine())
	time.Sleep(15 * time.Second)

	fmt.Println("======================main======================")
}

func TestClientGet(t *testing.T) {
	rawurl := "http://localhost:8080/n/v1/test?fruit=mango"
	httpClient := httpclient.NewClient(&httpclient.Config{})
	resp, err := httpClient.NewRequest(http.MethodGet, rawurl, nil).SetHeader(map[string]any{
		//"content-type": "application/x-www-form-urlencoded",
		//"content-type": "application/json",
		//"content-type": "application/octet-stream",
		"proxy_id":  100,
		"server_id": 2,
		"client_ip": "192.168.1.48",
		"user_id":   21,
		"trace_id":  15821793512,
	}).Do()
	fmt.Println("err:::::", err)
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
	httpClient := httpclient.NewClient(&httpclient.Config{})
	resp, err := httpClient.Post(rawurl, bodyByte).SetHeader(map[string]any{
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

func TestClientHead(t *testing.T) {
	rawurl := "http://localhost:8080/n/v1/test/post?sex=3"
	httpClient := httpclient.NewClient(&httpclient.Config{})
	resp, err := httpClient.NewRequest(http.MethodHead, rawurl, nil).SetHeader(map[string]any{
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
	fmt.Println("err:::::::::::", err)
	fmt.Println("resp.data:::::", string(data))

	fmt.Println("main")
}

func TestClientRequest(t *testing.T) {
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

	rawurl := "http://localhost:8080/n/v1/test/post?sex=3&age=18&test=x1"
	httpClient := httpclient.NewClient(&httpclient.Config{})
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

func TestClientRequestHttps(t *testing.T) {
	rawurl := "https://blog.csdn.net/Liing0/article/details/111241280?sex=3&age=18&test=x1"
	httpClient := httpclient.NewClient(&httpclient.Config{})
	resp, err := httpClient.Get(rawurl, map[string]any{"color": "red", "name": "tutu"}).Do()
	if err != nil {
		log.Fatal("err:::::::", err)
	}
	data, err := resp.GetData()
	fmt.Println("err:::::", err)
	fmt.Println("resp:::::", string(data))

	fmt.Println("main")
}
