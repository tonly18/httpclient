package httpclient

import (
	"errors"
	"net/http"
)

type HttpResponse struct {
	Response *http.Response
	Data     []byte
	Close    bool
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

func (r *HttpResponse) GetData() ([]byte, error) {
	defer r.Response.Body.Close()
	r.Response.Body.Close()
	if r.Close {
		return r.Data, nil
	}
	return nil, errors.New("response data happen error")
}
