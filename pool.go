package httpclient

import (
	"bytes"
	"sync"
)

var poolBuffer512K = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(make([]byte, 0, 1024*512))
	},
}

var poolBuffer1M = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(make([]byte, 0, 1024*1024))
	},
}

func poolGet(size string) *bytes.Buffer {
	if size == "512K" {
		buffer := poolBuffer512K.Get().(*bytes.Buffer)
		buffer.Reset()
		return buffer
	}
	if size == "1M" {
		buffer := poolBuffer1M.Get().(*bytes.Buffer)
		buffer.Reset()
		return buffer
	}

	return nil
}

func poolPut(size string, buffer *bytes.Buffer) {
	if size == "512K" {
		poolBuffer512K.Put(buffer)
	}
	if size == "1M" {
		poolBuffer1M.Put(buffer)
	}
}
