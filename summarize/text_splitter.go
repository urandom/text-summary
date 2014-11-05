package summarize

import "unicode"

type TextSplitter interface {
	Sentences(string) []string
	Words(string) []string
}

type DefaultTextSplitter struct {
	Punctuations []rune
}

func (d DefaultTextSplitter) Sentences(text string) []string {
	buf := getBuffer()
	defer bufferPool.Put(buf)

	sentences := []string{}
	newSentence := true
	lastNonWhiteSpace := -1

	for _, r := range text {
		if oneOfPunct(r, d.Punctuations) {
			if buf.Len() > 0 {
				if lastNonWhiteSpace > 0 {
					buf.Truncate(lastNonWhiteSpace)
					buf.WriteRune(r)
					sentences = append(sentences, buf.String())
				}
				buf.Reset()
				newSentence = true
			}
		} else {
			isSpace := unicode.IsSpace(r)
			if newSentence && isSpace {
				continue
			}
			newSentence = false
			buf.WriteRune(r)
			if !isSpace {
				lastNonWhiteSpace = buf.Len()
			}
		}
	}

	if buf.Len() > 0 && lastNonWhiteSpace > 0 {
		buf.Truncate(lastNonWhiteSpace)
		sentences = append(sentences, buf.String())
		buf.Reset()
	}

	return sentences
}

func (d DefaultTextSplitter) Words(text string) []string {
	buf := getBuffer()
	defer bufferPool.Put(buf)

	words := []string{}

	for _, r := range text {
		if unicode.IsLetter(r) {
			buf.WriteRune(r)
		} else {
			if buf.Len() > 0 {
				words = append(words, buf.String())
			}
			buf.Reset()
		}
	}

	if buf.Len() > 0 {
		words = append(words, buf.String())
	}

	return words
}

func oneOfPunct(r rune, punct []rune) bool {
	for _, p := range punct {
		if p == r {
			return true
		}
	}
	return false
}

func oneOfStartQuote(r rune, quotes [][]rune) int {
	for i, q := range quotes {
		if q[0] == r {
			return i
		}
	}
	return -1
}
