package httpclient

import (
	"bytes"
	"sync"
)

const maxSize = 1 << 16 //64KiB

var poolBuffer = &sync.Pool{
	New: func() any {
		return bytes.NewBuffer(make([]byte, 0, 512))
	},
}

func poolGet() *bytes.Buffer {
	return poolBuffer.Get().(*bytes.Buffer)
}

func poolPut(buffer *bytes.Buffer) {
	buffer.Reset()
	if buffer.Cap() > maxSize {
		return
	}
	poolBuffer.Put(buffer)
}
