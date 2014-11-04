package summarize

import "unicode"

type SentenceSplitter interface {
	Sentences(string) []string
}

type DefaultSentenceSplitter struct {
	Punctuations []rune
	Quotes       [][]rune
}

func (d DefaultSentenceSplitter) Sentences(text string) []string {
	buf := getBuffer()
	defer bufferPool.Put(buf)

	sentences := []string{}
	startQuote := -1
	newSentence := true
	lastNonWhiteSpace := -1

	for _, r := range text {
		if oneOfPunct(r, d.Punctuations) {
			if buf.Len() > 0 {
				if lastNonWhiteSpace > 0 {
					sentences = append(sentences, buf.String()[:lastNonWhiteSpace])
				}
				buf.Reset()
				newSentence = true
			}
		} else if startQuote > -1 && d.Quotes[startQuote][1] == r {
			if lastNonWhiteSpace > 0 {
				sentences = append(sentences, buf.String()[:lastNonWhiteSpace])
			}
			buf.Reset()
			newSentence = true
			startQuote = -1
		} else if startQuote = oneOfStartQuote(r, d.Quotes); startQuote > -1 {
			if buf.Len() > 0 {
				if lastNonWhiteSpace > 0 {
					sentences = append(sentences, buf.String()[:lastNonWhiteSpace])
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
		sentences = append(sentences, buf.String()[:lastNonWhiteSpace])
		buf.Reset()
	}

	return sentences
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
