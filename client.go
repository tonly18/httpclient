package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var client *http.Client
var once sync.Once

type HttpClient struct {
	httpClient   *http.Client
	httpRequest  *Request
	responseSize string //返回值大小:512K、1M、2M、5M
}

func NewHttpClient(config *Config) *HttpClient {
	if config == nil {
		config = &Config{}
	}
	if config.Transport == nil {
		config.Transport = transport
	}
	if config.TimeOut == 0 {
		config.TimeOut = time.Second * defaultTimeout //请求超时: 默认5秒
	}
	if config.ResponseSize == "" {
		config.ResponseSize = defaultSize //返回值大小: 默认1M
	}

	once.Do(func() {
		client = &http.Client{
			Transport:     config.Transport,
			CheckRedirect: config.CheckRedirect,
			Jar:           config.Jar,
			Timeout:       config.TimeOut, //从连接(Dial)到读完response body
		}
	})

	//return
	return &HttpClient{
		httpClient:   client,
		httpRequest:  nil,
		responseSize: config.ResponseSize,
	}
}

func (c *HttpClient) Get(rawurl string, params map[string]any) *HttpClient {
	if len(params) > 0 {
		urlValue := url.Values{}
		for k, v := range params {
			urlValue.Set(k, fmt.Sprintf(`%v`, v))
		}
		rawurl = fmt.Sprintf(`%v?%v`, rawurl, urlValue.Encode())
	}
	req, err := NewRequest(http.MethodGet, rawurl, nil)
	if err != nil {
		return nil
	}
	c.httpRequest = req

	return c
}

func (c *HttpClient) Post(rawurl string, params map[string]any) *HttpClient {
	req, err := NewRequest(http.MethodPost, rawurl, params)
	if err != nil {
		return nil
	}
	c.httpRequest = req

	return c
}

func (c *HttpClient) NewRequest(method, url string, params map[string]any) *HttpClient {
	if req, err := NewRequest(strings.ToUpper(method), url, params); err != nil {
		return nil
	} else {
		c.httpRequest = req
	}

	return c
}

func (c *HttpClient) SetHeader(params map[string]any) *HttpClient {
	c.httpRequest.SetHeader(params)

	//return
	return c
}

func (c *HttpClient) Do() (*HttpResponse, error) {
	resp, err := c.httpClient.Do(c.httpRequest.Request)
	c.httpRequest.Request.Body.Close()
	if err != nil {
		return nil, err
	}

	//response
	rawBuffer := poolGet(c.responseSize)
	defer func() {
		poolPut(c.responseSize, rawBuffer)
		resp.Body.Close()
	}()
	if _, err := io.Copy(rawBuffer, resp.Body); err != nil {
		return nil, err
	}

	//return
	return &HttpResponse{
		Response: resp,
		Data:     rawBuffer.Bytes(),
		Close:    true,
	}, nil
}
