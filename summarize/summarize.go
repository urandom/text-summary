package summarize

import "io"

type Summarize struct {
	Title             string
	Text              string
	Language          string
	StopWordsProvider StopWordsProvider
	TextSplitter      TextSplitter
	EndSentenceRunes  []rune
	QuoteTuples       [][]rune
}

func New(title string, r io.Reader) Summarize {
	buffer := getBuffer()
	defer bufferPool.Put(buffer)

	buffer.ReadFrom(r)

	return NewFromString(title, buffer.String())
}

func NewFromString(title, text string) Summarize {
	return Summarize{
		Title:             title,
		Text:              text,
		Language:          "en",
		StopWordsProvider: DefaultStopWords{},
		TextSplitter:      DefaultTextSplitter{[]rune{'.', '!', '?'}, [][]rune{[]rune{'\'', '\''}, []rune{'"', '"'}, []rune{'`', '`'}}},
	}
}

func (s Summarize) KeyPoints() []string {
	//sentenced := s.TextSplitter.Sentences(s.Text)

	return []string{}
}
