package httpclient

import (
	"fmt"
	"net/url"
	"strings"
)

// queryEncode
func queryEncode(rawurl string, params map[string]any) string {
	if len(params) > 0 {
		args := url.Values{}
		if strings.Contains(rawurl, "?") {
			urlInfo := strings.Split(rawurl, "?")
			rawurl = urlInfo[0]
			query := urlInfo[1]
			args, _ = url.ParseQuery(query)
		}
		for k, v := range params {
			args.Set(k, fmt.Sprintf(`%v`, v))
		}
		rawurl = fmt.Sprintf(`%v?%v`, rawurl, args.Encode())
	}

	//return
	return rawurl
}
