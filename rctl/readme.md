## 利用 chan 进行并发控制, 代码如下：

ctl.go

```Golang
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

```
ctl_test.go

```Golang
package rctl

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_FuncCtl(t *testing.T) {
	rt := Rctl{
		Queue: make(chan int, 400),
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

```
利用带缓冲区的chan可以很方便的进行并发控制，比如本例子就控制了最多400个并发(顺便复习一下函数的传递)
运行的结果应该是：先打印一堆的："go to af", 再打印：100, 101, 102...
