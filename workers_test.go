package prose

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewWorkerPool(t *testing.T) {
	j := make(chan *Token)
	r := make(chan *Token)
	e := make(chan bool)
	var f tagFn
	f = func(arg1 []*Token) []*Token {
		return []*Token{}
	}
	p := NewWorkerPool(j, r, e, f)
	if p == nil {
		panic(errors.New("pool is empty"))
	}
}

func TestTag(t *testing.T) {
	tkList := []*Token{&Token{Text: "John"}, &Token{Text: "world"}, &Token{Text: "or"}, &Token{Text: "for"}}
	j := make(chan *Token, 4)
	r := make(chan *Token, 4)
	e := make(chan bool)
	var f tagFn
	f = func(arg1 []*Token) []*Token {
		return []*Token{}
	}
	p := NewWorkerPool(j, r, e, f)

	for _, v := range tkList {
		j <- v
	}
	close(j)

	p.RunTagAndWait()
	if len(r) != 4 {
		panic(fmt.Errorf("length of results: %d", len(r)))
	}
}
