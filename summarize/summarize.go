package summarize

import (
	"bytes"
	"io"
)

type Summarize struct {
	Title             string
	Text              string
	Language          string
	StopWordsProvider StopWordsProvider
}

func New(title string, r io.Reader) Summarize {
	buffer := bufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer bufferPool.Put(buffer)

	buffer.ReadFrom(r)

	return NewFromString(title, buffer.String())
}

func NewFromString(title, text string) Summarize {
	return Summarize{
		Title:             title,
		Text:              text,
		Language:          "en",
		StopWordsProvider: defaultStopWords{},
	}
}
