package summarize

import "testing"

func TestDefaultSentenceSplitter(t *testing.T) {
	d := DefaultTextSplitter{[]rune{'.', '!', '?'}, [][]rune{[]rune{'\'', '\''}, []rune{'"', '"'}, []rune{'`', '`'}}}

	exp := []string{"First sentence", "Second sentence"}
	res := d.Sentences(text1)

	checkSentences(t, exp, res)

	exp = []string{"First sentence"}
	res = d.Sentences(text2)

	checkSentences(t, exp, res)

	exp = []string{"Then he said:", "do not feed the sharks"}
	res = d.Sentences(text3)

	checkSentences(t, exp, res)

	exp = []string{
		"The contribution of cloud computing and mobile computing technologies lead to the newly emerging mobile cloud com- puting paradigm",
		"Three major approaches have been pro- posed for mobile cloud applications: 1) extending the access to cloud services to mobile devices; 2) enabling mobile de- vices to work collaboratively as cloud resource providers; 3) augmenting the execution of mobile applications on portable devices using cloud resources",
		"In this paper, we focus on the third approach in supporting mobile data stream applica- tions",
		"More specifically, we study how to optimize the com- putation partitioning of a data stream application between mobile and cloud to achieve maximum speed/throughput in processing the streaming data",
		"To the best of our knowledge, it is the first work to study the partitioning problem for mobile data stream applica- tions, where the optimization is placed on achieving high throughput of processing the streaming data rather than minimizing the makespan of executions as in other appli- cations",
	}
	res = d.Sentences(bigText)

	checkSentences(t, exp, res)
}

func TestDefaultWordSplitter(t *testing.T) {
	d := DefaultTextSplitter{[]rune{'.', '!', '?'}, [][]rune{[]rune{'\'', '\''}, []rune{'"', '"'}, []rune{'`', '`'}}}

	exp := []string{"First", "sentence", "Second", "sentence"}
	res := d.Words(text1)

	checkSentences(t, exp, res)

	exp = []string{"Then", "he", "said", "do", "not", "feed", "the", "sharks"}
	res = d.Words(text3)

	checkSentences(t, exp, res)
}

func checkSentences(t *testing.T, exp, res []string) {
	if len(exp) != len(res) {
		t.Fatalf("Number of sentences differ from expected: %d - %d\n", len(res), len(exp))
	}
	for i, s := range exp {
		if s != res[i] {
			t.Fatalf("Expected sentence '%s', got '%s'\n", s, res[i])
		}
	}
}

var (
	text1   = ` First sentence.   Second sentence  `
	text2   = `First sentence.    `
	text3   = `Then he said: " do not feed the sharks"`
	bigText = `The contribution of cloud computing and mobile computing technologies lead to the newly emerging mobile cloud com- puting paradigm.
Three major approaches have been pro- posed for mobile cloud applications: 1) extending the access to cloud services to mobile devices; 2) enabling mobile de- vices to work collaboratively as cloud resource providers; 3) augmenting the execution of mobile applications on portable devices using cloud resources.
In this paper, we focus on the third approach in supporting mobile data stream applica- tions.
More specifically, we study how to optimize the com- putation partitioning of a data stream application between mobile and cloud to achieve maximum speed/throughput in processing the streaming data.
To the best of our knowledge, it is the first work to study the partitioning problem for mobile data stream applica- tions, where the optimization is placed on achieving high throughput of processing the streaming data rather than minimizing the makespan of executions as in other appli- cations.`
)
