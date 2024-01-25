package httpclient

import (
	"net/http"
	"sync"
	"time"
)

var once sync.Once
var rawclient *rawClient

type rawClient struct {
	client *http.Client
}

func NewClient(config *Config) *rawClient {
	if rawclient != nil {
		return rawclient
	}

	if config == nil {
		config = &Config{}
	}
	if config.Transport == nil {
		config.Transport = transport
	}
	if config.TimeOut == 0 {
		config.TimeOut = time.Second * httpRequestTimeOut
	}

	once.Do(func() {
		rawclient = &rawClient{
			client: &http.Client{
				Transport:     config.Transport,
				CheckRedirect: config.CheckRedirect,
				Jar:           config.Jar,
				Timeout:       config.TimeOut, //从连接(Dial)到读完response body
			},
		}
	})

	//return
	return rawclient
}

// GET请求
func (c *rawClient) Get(rawurl string, params map[string]any) *httpClient {
	rawurl = queryEncode(rawurl, params)
	req, err := NewRequest(http.MethodGet, rawurl, nil)
	if err != nil {
		return nil
	}

	//return
	return &httpClient{
		httpClient:  c.client,
		httpRequest: req,
	}
}

// POST请求
func (c *rawClient) Post(rawurl string, body []byte) *httpClient {
	req, err := NewRequest(http.MethodPost, rawurl, body)
	if err != nil {
		return nil
	}

	//return
	return &httpClient{
		httpClient:  c.client,
		httpRequest: req,
	}
}

// http请求
func (c *rawClient) NewRequest(method, rawurl string, body []byte) *httpClient {
	req, err := NewRequest(method, rawurl, body)
	if err != nil {
		return nil
	}

	//return
	return &httpClient{
		httpClient:  c.client,
		httpRequest: req,
	}
}
