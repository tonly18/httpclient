package httpclient

import (
	"fmt"
	"net/url"
)

// queryEncode
func queryEncode(rawurl string, params map[string]any) string {
	if len(params) > 0 {
		urlArgs := url.Values{}
		for k, v := range params {
			urlArgs.Set(k, fmt.Sprintf(`%v`, v))
		}
		rawurl = fmt.Sprintf(`%v?%v`, rawurl, urlArgs.Encode())
	}

	//return
	return rawurl
}
