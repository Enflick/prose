package prose

import (
	"path/filepath"
	"testing"
)

func BenchmarkDoc(b *testing.B) {
	content := readDataFile(filepath.Join(testdata, "sherlock.txt"))
	text := string(content)
	for n := 0; n < b.N; n++ {
		_, err := NewDocument(text)
		if err != nil {
			panic(err)
		}
	}
}

func TestConcurrentDocumentTag(t *testing.T) {
	text := string("they'll won't go if they are not bad")
	_, err := NewDocument(text, WithExtraction(false), WithConcurrency(true))
	if err != nil {
		panic(err)
	}
}
