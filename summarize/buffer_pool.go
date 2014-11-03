package summarize

import (
	"bytes"
	"sync"
)

var bufferPool sync.Pool

func init() {
	bufferPool = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
}
