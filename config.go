package httpclient

import (
	"crypto/tls"
	"net/http"
	"time"
)

const (
	defaultTimeout = 15
	contentType    = "Content-Type"
)

// config struct
type Config struct {
	Transport     *http.Transport
	CheckRedirect func(req *http.Request, via []*http.Request) error
	Jar           http.CookieJar
	TimeOut       time.Duration
}

// http.Transport内都会维护一个自己的空闲连接池,如果每个client都创建一个新的http.Transport,就会导致底层的TCP连接无法复用.
// 如果网络请求过大,上面这种情况会导致协程数量变得非常多,导致服务不稳定.
var transport = &http.Transport{
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true, //不校验服务端证书
	},
	DisableKeepAlives:     false,            //禁用keep-alive
	MaxConnsPerHost:       2000,             //每个Host(IP:PORT元组)能创建的最大连接数(默认值:0)
	MaxIdleConns:          100,              //最大空闲数量(默认值:100)
	MaxIdleConnsPerHost:   1000,             //每个Host最大空闲数(能复用的链接数)(默认值:2)
	IdleConnTimeout:       90 * time.Second, //连接空闲超时(默认值:90)
	DisableCompression:    true,             //禁止压缩
	ExpectContinueTimeout: 1 * time.Second,  //(默认值:1)限制client在发送包含Expect:100-continue的header到收到继续发送body的response之间的时间等待

	//ResponseHeaderTimeout: 10 * time.Second,	//从client.request执行do之后,到server返回截至
	//ForceAttemptHTTP2: true,             		//强制使用http2(默认值:true)
	//TLSHandshakeTimeout:   10 * time.Second,  //TLS握手超时(默认值:10)
}
