package httpclient

import (
	"fmt"
	"net/url"
)

// queryEncode
func queryEncode(rawurl string, params map[string]any) string {
	if len(params) > 0 {
		urlInfo, _ := url.Parse(rawurl)
		query := urlInfo.Query()
		for k, v := range params {
			query.Set(k, fmt.Sprintf(`%v`, v))
		}
		urlInfo.RawQuery = query.Encode()
		rawurl = urlInfo.String()
	}

	//return
	return rawurl
}
