package httpclient

import (
	"errors"
	"net/http"
)

type HttpResponse struct {
	Response *http.Response
	Data     []byte
	Close    bool
	Error    error
}

func NewHttpResponse(response *http.Response) *HttpResponse {
	return &HttpResponse{
		Response: response,
	}
}

func (r *HttpResponse) GetHeaderCode() int {
	if r.Response != nil {
		return r.Response.StatusCode
	}
	return 0
}

func (r *HttpResponse) GetDataFromHeader(key string) string {
	if r.Response != nil {
		return r.Response.Header.Get(key)
	}
	return ""
}

func (r *HttpResponse) GetCookie(names ...string) map[string]*http.Cookie {
	data := make(map[string]*http.Cookie, len(names))
	for _, cname := range names {
		for _, v := range r.Response.Cookies() {
			if v.Name == cname {
				data[cname] = v
				break
			}
		}
	}

	return data
}

func (r *HttpResponse) GetData() ([]byte, error) {
	defer func() {
		if r.Response != nil && r.Response.Body != nil {
			r.Response.Body.Close()
		}
	}()

	if r.Response == nil {
		return nil, errors.New("response is nil")
	}

	if r.Close {
		return r.Data, nil
	}

	return nil, errors.New("get response data happen error")
}
