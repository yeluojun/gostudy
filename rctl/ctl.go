package rctl

import "sync"

type Rctl struct {
	Queue chan int
	Wg    *sync.WaitGroup
}

func (r *Rctl) FuncCtl(f func(int), i int) {
	// 并发操作
	go func() {
		r.Queue <- 1
		f(i)
		<-r.Queue
		r.Wg.Done()
	}()
}
