package summarize

import (
	"io"
	"math"
)

type Summarize struct {
	Title             string
	Text              string
	Language          string
	StopWordsProvider StopWordsProvider
	TextSplitter      TextSplitter
	EndSentenceRunes  []rune
	QuoteTuples       [][]rune
	IdealWordCount    int
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
		IdealWordCount:    20,
	}
}

func (s Summarize) KeyPoints() []string {
	sentences := s.TextSplitter.Sentences(s.Text)
	keywords := s.keywords(s.Text)
	titleWords := s.TextSplitter.Words(s.Title)

	if len(sentences) <= 5 {
		return sentences
	}

	ranks := TextCounter{}

	titleMap := map[string]bool{}
	for _, t := range titleWords {
		if !s.StopWordsProvider.IsStopWord(t) {
			titleMap[t] = true
		}
	}

	for i, sent := range sentences {
		words := s.TextSplitter.Words(sent)
		titleScore := s.titleScore(titleMap, words)
		lengthScore := s.lengthScore(words)
		positionScore := s.positionScore(i, len(sentences))
		sbs := s.sbs(words, keywords)
		dbs := s.dbs(words, keywords)

		freq := (sbs + dbs) / 2 * 10
		total := (titleScore*1.5 + freq*2 + lengthScore + positionScore) / 4
		ranks.AddScored(sent, int(total*100))
	}

	var keyPoints []string

	mostCommon := ranks.MostCommon(5)

	for _, p := range mostCommon {
		keyPoints = append(keyPoints, p.Text)
	}

	return keyPoints
}

func (s Summarize) keywords(text string) map[string]float64 {
	allWords := s.TextSplitter.Words(text)
	allLen := float64(len(allWords))
	filteredWords := []string{}

	for _, w := range allWords {
		if !s.StopWordsProvider.IsStopWord(w) {
			filteredWords = append(filteredWords, w)
		}
	}

	freq := NewTextCounterFromSlice(filteredWords)

	pairs := freq.MostCommon(10)
	keyMap := map[string]float64{}

	for _, p := range pairs {
		score := float64(p.Count) / allLen
		keyMap[p.Text] = score * 15
	}

	return keyMap
}

func (s Summarize) titleScore(titleMap map[string]bool, words []string) float64 {
	count := 0

	for _, w := range words {
		if _, ok := titleMap[w]; ok && !s.StopWordsProvider.IsStopWord(w) {
			count += 1
		}
	}

	return float64(count) / float64(len(words))
}

func (s Summarize) lengthScore(words []string) float64 {
	return 1 - (math.Abs(float64(s.IdealWordCount-len(words))) / float64(s.IdealWordCount))
}

func (s Summarize) positionScore(pos, total int) float64 {
	normalized := float64(pos) / float64(total)

	if normalized < 0 {
		normalized = 0
	}

	if normalized <= 0.1 {
		return 0.17
	} else if normalized <= 0.2 {
		return 0.23
	} else if normalized <= 0.3 {
		return 0.14
	} else if normalized <= 0.4 {
		return 0.08
	} else if normalized <= 0.5 {
		return 0.05
	} else if normalized <= 0.6 {
		return 0.04
	} else if normalized <= 0.7 {
		return 0.06
	} else if normalized <= 0.8 {
		return 0.04
	} else if normalized <= 0.9 {
		return 0.04
	} else if normalized <= 1.0 {
		return 0.15
	} else {
		return 0
	}
}

func (s Summarize) sbs(words []string, keywords map[string]float64) float64 {
	score := 0.0

	if len(words) == 0 {
		return score
	}

	for _, w := range words {
		if c, ok := keywords[w]; ok {
			score += c
		}
	}

	return (1 / float64(len(words)) * score) / 10
}

func (s Summarize) dbs(words []string, keywords map[string]float64) float64 {
	score := 0.0

	if len(words) == 0 {
		return score
	}

	summ := 0.0
	first := [2]float64{}
	second := [2]float64{}

	uniqueWords := map[string]bool{}
	for i, w := range words {
		if c, ok := keywords[w]; ok {
			if len(first) == 0 {
				first[0], first[1] = float64(i), c
			} else {
				second[0], second[1] = first[0], first[1]
				first[0], first[1] = float64(i), c

				diff := first[0] - second[0]
				summ += first[1] * second[1] / math.Pow(diff, 2)
			}

			uniqueWords[w] = true
		}
	}

	k := float64(len(uniqueWords) + 1)
	return (1 / (k * (k + 1)) * summ)
}
