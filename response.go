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
	return r.Response.StatusCode
}

func (r *HttpResponse) GetDataFromHeader(key string) string {
	return r.Response.Header.Get(key)
}

func (r *HttpResponse) Cookie(name string) *http.Cookie {
	for _, v := range r.Response.Cookies() {
		if v.Name == name {
			return v
		}
	}
	return nil
}

func (r *HttpResponse) GetData() ([]byte, error) {
	defer func() {
		if r.Response.Body != nil {
			r.Response.Body.Close()
		}
	}()

	if r.Close {
		return r.Data, nil
	}

	return nil, errors.New("response data happen error")
}
