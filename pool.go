package httpclient

import (
	"bytes"
	"sync"
)

var buffer1024Pool = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(make([]byte, 0, 1024))
	},
}
