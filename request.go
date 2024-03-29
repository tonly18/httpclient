package httpclient

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptrace"
	"strconv"
	"strings"
	"time"
)

type HttpRequest struct {
	Request *http.Request
	Method  string
}

func NewRequest(method, rawurl string, body []byte, debug bool) (*HttpRequest, error) {
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
	if len(body) > 0 {
		req.Header.Set("Content-Length", strconv.Itoa(len(body)))
	}

	//trace
	if debug {
		trace := &httptrace.ClientTrace{
			DNSStart: func(info httptrace.DNSStartInfo) {
				nowtime := time.Now().String()
				fmt.Println(nowtime, "dns start")
			},
			DNSDone: func(info httptrace.DNSDoneInfo) {
				nowtime := time.Now().String()
				fmt.Println(nowtime, "dns end")
			},
			ConnectStart: func(network, addr string) {
				nowtime := time.Now().String()
				fmt.Println(nowtime, "dial start")
			},
			ConnectDone: func(network, addr string, err error) {
				nowtime := time.Now().String()
				fmt.Println(nowtime, "dial end")
			},
			GotConn: func(connInfo httptrace.GotConnInfo) {
				nowtime := time.Now().String()
				fmt.Println(nowtime, "conn time", connInfo)
			},
			WroteHeaders: func() {
				nowtime := time.Now().String()
				fmt.Println(nowtime, "wrote all request headers")
			},
			WroteRequest: func(wr httptrace.WroteRequestInfo) {
				nowtime := time.Now().String()
				fmt.Println(nowtime, "wrote all request")
			},
			GotFirstResponseByte: func() {
				nowtime := time.Now().String()
				fmt.Println(nowtime, "first received response byte")
			},
		}
		req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
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
		if strings.ToLower(key) == strings.ToLower(contentType) {
			key = contentType
		}
		r.Request.Header.Set(key, fmt.Sprintf(`%v`, value))
	}

	return r
}

// AddCookie 设置cookie
func (r *HttpRequest) AddCookie(cookies []*http.Cookie) *HttpRequest {
	for _, v := range cookies {
		r.Request.AddCookie(v)
	}

	return r
}

// Prepare 准备request
func (r *HttpRequest) Prepare() *http.Request {
	if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodPatch {
		if r.Request.Header.Get(contentType) == "" {
			r.Request.Header.Set(contentType, "text/plain; charset=utf-8")
		}
	}

	return r.Request
}
