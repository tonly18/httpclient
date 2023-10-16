package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Request *http.Request
}

func NewRequest(method, url string, params map[string]any) (*Request, error) {
	var msgbody []byte
	if len(params) > 0 {
		msgbody, _ = json.Marshal(params)
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(msgbody))
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
		r.Request.Header.Set(key, fmt.Sprintf(`%v`, value))
	}

	return r
}
