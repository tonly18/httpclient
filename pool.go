package httpclient

import (
	"bytes"
	"sync"
)

const maxSize = 1 << 16 //64KiB

var poolBuffer = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(make([]byte, 0, 512))
	},
}

func poolGet() *bytes.Buffer {
	buffer := poolBuffer.Get().(*bytes.Buffer)
	buffer.Reset()

	return buffer
}

func poolPut(buffer *bytes.Buffer) {
	if buffer.Cap() > maxSize {
		return
	}
	poolBuffer.Put(buffer)
}
