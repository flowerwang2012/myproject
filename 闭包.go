package main

import (
	"container/list"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg2 sync.WaitGroup
	wg2.Add(12)
	for i := 0; i < 6; i++ {
		go func() {
			fmt.Println("T1: ", i)
			wg2.Done()
		}()
	}
	for i := 0; i < 6; i++ {
		go func(i int) {
			fmt.Println("T2: ", i)
			wg2.Done()
		}(i)
	}
	wg2.Wait()
	// 控制goroutine的执行顺序，用双向列表list控制
	wg2.Add(6)
	l := list.New()
	for i := 0; i < 6; i++ {
		l.PushBack(i)
	}
	for i := 0; i < 6; i++ {
		go func(i int) {
			for {
				elem := l.Front()
				n := elem.Value.(int)
				if i == n {
					fmt.Println("T3: ", i)
					l.Remove(elem)
					wg2.Done()
					return
				} else {
					time.Sleep(50 * time.Millisecond)
				}
			}
		}(i)
	}
	wg2.Wait()
}
