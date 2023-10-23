package httpclient

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

type HttpRequest struct {
	Request *http.Request
	Method  string
}

func NewRequest(method, rawurl string, body []byte) (*HttpRequest, error) {
	var err error
	var req *http.Request

	if len(body) == 0 {
		req, err = http.NewRequest(method, rawurl, nil)
	} else {
		req, err = http.NewRequest(method, rawurl, bytes.NewBuffer(body))
	}
	if err != nil {
		return nil, err
	}

	//return
	return &HttpRequest{
		Request: req,
		Method:  method,
	}, nil
}

// SetHeader 设置头信息
func (r *HttpRequest) SetHeader(params map[string]any) *HttpRequest {
	for key, value := range params {
		if strings.Index(strings.ToLower(key), strings.ToLower(contentType)) > -1 {
			key = contentType
		}
		r.Request.Header.Set(key, fmt.Sprintf(`%v`, value))
	}

	return r
}

// Prepare 准备request
func (r *HttpRequest) Prepare() *http.Request {
	if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodPatch {
		if r.Request.Header.Get(contentType) == "" {
			r.Request.Header.Set(contentType, "application/x-www-form-urlencoded")
		}
	}

	return r.Request
}
