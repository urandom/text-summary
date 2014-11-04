package summarize

import (
	"bytes"
	"sync"
)

var bufferPool sync.Pool

func getBuffer() *bytes.Buffer {
	buffer := bufferPool.Get().(*bytes.Buffer)
	buffer.Reset()

	return buffer
}

func init() {
	bufferPool = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
}
