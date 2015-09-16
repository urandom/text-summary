package summarize

import "sort"

type TextCounter map[string]int
type CommonPairs []CommonPair
type CommonPair struct {
	Text  string
	Count int
}

func NewTextCounterFromPairs(pairs CommonPairs) TextCounter {
	tc := TextCounter{}
	for _, p := range pairs {
		tc[p.Text] = p.Count
	}

	return tc
}

func NewTextCounterFromSlice(words []string) TextCounter {
	tc := TextCounter{}
	for _, w := range words {
		tc.Add(w)
	}

	return tc
}

func (tc TextCounter) Add(text string, score ...int) {
	sc := 1
	if len(score) > 0 {
		sc = score[0]
	}

	if c, ok := tc[text]; ok {
		tc[text] = c + sc
	} else {
		tc[text] = sc
	}
}

func (tc TextCounter) MostCommon(limit ...int) CommonPairs {
	pairs := CommonPairs{}

	for t, c := range tc {
		pairs = append(pairs, CommonPair{Text: t, Count: c})
	}

	sort.Sort(pairs)

	if len(limit) > 0 && len(pairs) > limit[0] {
		return pairs[:limit[0]]
	}

	return pairs
}

func (c CommonPairs) Len() int {
	return len(c)
}

func (c CommonPairs) Less(i, j int) bool {
	if c[i].Count == c[j].Count {
		return len(c[i].Text) > len(c[j].Text)
	}
	return c[i].Count > c[j].Count
}

func (c CommonPairs) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
