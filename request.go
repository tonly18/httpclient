package httpclient

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

type Request struct {
	Request *http.Request
	Method  string
}

func NewRequest(method, rawurl string, body []byte) (*Request, error) {
	var err error
	var req *http.Request

	if body == nil {
		req, err = http.NewRequest(method, rawurl, nil)
	} else {
		req, err = http.NewRequest(method, rawurl, bytes.NewBuffer(body))
	}
	if err != nil {
		return nil, err
	}

	//return
	return &Request{
		Request: req,
		Method:  method,
	}, nil
}

// SetHeader 设置头信息
func (r *Request) SetHeader(params map[string]any) *Request {
	for key, value := range params {
		if strings.Index(strings.ToLower(key), contentType) > -1 {
			key = contentType
		}
		r.Request.Header.Set(key, fmt.Sprintf(`%v`, value))
	}

	return r
}

// Prepare 准备request
func (r *Request) Prepare() *http.Request {
	if r.Method != http.MethodGet {
		if r.Request.Header.Get(contentType) == "" {
			r.Request.Header.Set(contentType, "application/x-www-form-urlencoded")
		}
	}

	return r.Request
}
