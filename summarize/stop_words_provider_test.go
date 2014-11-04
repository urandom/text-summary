package summarize

import "testing"

func TestDefaultStopWords(t *testing.T) {
	d := DefaultStopWords{}

	if !d.IsStopWord("a") {
		t.Fatal("Expected 'a' to be a stop word")
	}
}
