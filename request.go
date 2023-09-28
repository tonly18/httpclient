package httpclient

import (
	"bytes"
	"github.com/spf13/cast"
	"net/http"
)

type Request struct {
	Request *http.Request
}

func NewRequest(method, url string, data []byte) (*Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	//return
	return &Request{
		Request: req,
	}, nil
}

// SetHeader 设置头信息
func (r *Request) SetHeader(params map[string]any) *Request {
	for key, value := range params {
		r.Request.Header.Set(key, cast.ToString(value))
	}

	return r
}
