package httpclient

import (
	"io"
	"net/http"
)

type httpClient struct {
	httpClient   *http.Client
	httpRequest  *HttpRequest
	httpResponse *HttpResponse
}

func (c *httpClient) SetHeader(params map[string]any) *httpClient {
	c.httpRequest.SetHeader(params)

	//return
	return c
}

func (c *httpClient) AddCookie(cookies []*http.Cookie) *httpClient {
	c.httpRequest.AddCookie(cookies)

	//return
	return c
}

func (c *httpClient) Do() (*HttpResponse, error) {
	resp := c.DoNew()
	return resp, resp.Error
}

func (c *httpClient) DoNew() *HttpResponse {
	resp, err := c.httpClient.Do(c.httpRequest.Prepare())
	if c.httpRequest.Request.Body != nil {
		c.httpRequest.Request.Body.Close()
	}
	if err != nil {
		return &HttpResponse{
			Error: err,
		}
	}

	//response
	rawBuffer := poolGet()
	defer func() {
		poolPut(rawBuffer)
		_ = resp.Body.Close()
	}()
	if _, err := io.Copy(rawBuffer, resp.Body); err != nil {
		return &HttpResponse{
			Error: err,
		}
	}

	//return
	return &HttpResponse{
		Response: resp,
		Data:     rawBuffer.Bytes(),
		Close:    true,
	}
}
