package pool

import "sync"

type Pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

func NewPool(size int) *Pool {
	if size <= 0 {
		size = 1
	}
	return &Pool{queue: make(chan int, size), wg: &sync.WaitGroup{}}
}

func (this *Pool) Add(size int) {
	for i := 0; i < size; i++ { // size > 0
		this.queue <- 1
	}
	for i := 0; i > size; i-- { // size < 0
		<-this.queue
	}
	this.wg.Add(size)
}

func (this *Pool) Done() {
	<-this.queue
	this.wg.Done()
}

func (this *Pool) Wait() {
	this.wg.Wait()
}
