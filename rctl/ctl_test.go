package rctl

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_FuncCtl(t *testing.T) {
	rt := Rctl{
		Queue: make(chan int, 300),
		Wg:    new(sync.WaitGroup),
	}

	rt.Wg.Add(400)
	for i := 100; i < 500; i++ {
		rt.FuncCtl(Af, i)
	}
	rt.Wg.Wait()
	close(rt.Queue)
}

func Af(t int) {
	fmt.Println("go to af")
	time.Sleep(time.Millisecond * time.Duration(t))
	fmt.Println(t)
}
