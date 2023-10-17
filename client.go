package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var client *http.Client
var once sync.Once

type HttpClient struct {
	httpClient   *http.Client
	httpRequest  *HttpRequest
	responseSize string //返回值大小:512K、1M
}

func NewHttpClient(config *Config) *HttpClient {
	if config == nil {
		config = &Config{}
	}
	if config.Transport == nil {
		config.Transport = transport
	}
	if config.TimeOut == 0 {
		config.TimeOut = time.Second * defaultTimeout //请求超时: 默认15秒
	}
	if config.ResponseSize == "" {
		config.ResponseSize = defaultSize //返回值大小: 默认512K
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
		urlParams := url.Values{}
		for k, v := range params {
			urlParams.Set(k, fmt.Sprintf(`%v`, v))
		}
		rawurl = fmt.Sprintf(`%v?%v`, rawurl, urlParams.Encode())
	}
	req, err := NewRequest(http.MethodGet, rawurl, nil)
	if err != nil {
		return nil
	}
	c.httpRequest = req

	return c
}

func (c *HttpClient) Post(rawurl string, body []byte) *HttpClient {
	req, err := NewRequest(http.MethodPost, rawurl, body)
	if err != nil {
		return nil
	}
	c.httpRequest = req

	return c
}

func (c *HttpClient) NewRequest(method, rawurl string, body []byte) *HttpClient {
	if req, err := NewRequest(method, rawurl, body); err != nil {
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
	resp, err := c.httpClient.Do(c.httpRequest.Prepare())
	if c.httpRequest.Request.Body != nil {
		c.httpRequest.Request.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	//response
	rawBuffer := poolGet(c.responseSize)
	defer func() {
		resp.Body.Close()
		poolPut(c.responseSize, rawBuffer)
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
