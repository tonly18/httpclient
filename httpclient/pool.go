package httpclient

import (
	"bytes"
	"sync"
)

var buffer64Pool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

var buffer512Pool = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(make([]byte, 0, 512))
	},
}

var buffer1024Pool = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(make([]byte, 0, 1024))
	},
}

var buffer2408Pool = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(make([]byte, 0, 2408))
	},
}

var buffer4096Pool = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(make([]byte, 0, 4096))
	},
}
