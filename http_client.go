package httpclient

import (
	"io"
	"net/http"
)

type httpClient struct {
	httpClient  *http.Client
	httpRequest *HttpRequest
}

func (c *httpClient) SetHeader(params map[string]any) *httpClient {
	c.httpRequest.SetHeader(params)

	//return
	return c
}

func (c *httpClient) Do() (*HttpResponse, error) {
	resp, err := c.httpClient.Do(c.httpRequest.Prepare())
	if c.httpRequest.Request.Body != nil {
		c.httpRequest.Request.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	//response
	rawBuffer := poolGet()
	defer func() {
		poolPut(rawBuffer)
		_ = resp.Body.Close()
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
