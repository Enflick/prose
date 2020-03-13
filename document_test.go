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
	text := string("Mary had a little lamb Mary had a little lamb Mary had a little lamb")
	doc, err := NewDocument(text, WithExtraction(false), WithConcurrency(true))
	if err != nil {
		panic(err)
	}
	for _, i := range doc.tokens {
		if i.Tag == "" {
			panic("tag not done")
		}
	}
}

func BenchmarkConcurrentTagging(b *testing.B) {
	text := string("Mary had a little lamb Mary had a little lamb Mary had a little lamb Mary had a little lamb Mary had a little lamb")
	for n := 0; n < b.N; n++ {
		_, err := NewDocument(text, WithExtraction(false), WithConcurrency(true))
		if err != nil {
			panic(err)
		}
	}
}
