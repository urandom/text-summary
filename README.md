text-summary
============

Text-summary takes some text, and extracts its key points. Its an almost direct port of [PyTeaser](https://github.com/xiaoxu193/PyTeaser)

# From PyTeaser

Summaries are created by ranking sentences in a news article according to how relevant they are to the entire text. The top 5 sentences are used to form a "summary". Each sentence is ranked by using four criteria:

- Relevance to the title
- Relevance to keywords in the article
- Position of the sentence
- Length of the sentence

# Usage
```
import "github.com/urandom/text-summary/summarize"

...

s := summarize.New("Title for the text", someIOReader)
// or
// s := summarize.NewFromString("Title for the text", "Lengthy text ...")

keyPoints := s.KeyPoints()
```
