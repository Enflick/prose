package prose

import (
	"runtime"
	"sync"
)

// maxQueue is the buffer of the channel that workers get the items from
var maxQueue = 20

type tagFn func([]*Token) []*Token

// WorkerPool defines a pool of workers with the max number of workers
type WorkerPool struct {
	workers       int
	JobChannel    chan *Token
	ResultChannel chan *Token
	End           chan bool
	tagTk         tagFn
	wg            sync.WaitGroup
}

// NewWorkerPool returns a new pool of workers
func NewWorkerPool(jobs chan *Token, results chan *Token, end chan bool, tg tagFn) *WorkerPool {
	// numWorkers is the number of workers we run concurrently
	var maxWorkers = runtime.NumCPU()

	return &WorkerPool{workers: maxWorkers, JobChannel: jobs, ResultChannel: results, End: end, tagTk: tg}
}

// RunTagAndWait starts up the pool of workers and waits for completion
func (w *WorkerPool) RunTagAndWait() {
	for i := 0; i < w.workers; i++ {
		w.wg.Add(1)
		go w.tag()
	}
	w.wg.Wait()
}

// tag is used to actually tag the tokens
func (w *WorkerPool) tag() {
	for {
		j, ok := <-w.JobChannel
		if !ok {
			return
		}
		jt := []*Token{j}
		taggedTk := w.tagTk(jt)
		for _, item := range taggedTk {
			w.ResultChannel <- item
		}
	}
}
