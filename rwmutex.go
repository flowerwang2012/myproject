package main

import (
	"sync"
	"fmt"
	"math/rand"
)

var count1 int
var wg3 sync.WaitGroup
var rw sync.RWMutex //可以多读，但是只能有一个写

func main() {
	wg3.Add(10)

	for i:=0;i<5;i++ {
		go read(i)
	}

	for i:=0;i<5;i++ {
		go write(i);
	}

	wg3.Wait()
}

func read(n int) {
	rw.RLock()
	fmt.Printf("读goroutine %d 正在读取...\n",n)

	v := count1

	fmt.Printf("读goroutine %d 读取结束，值为：%d\n", n,v)
	wg3.Done()
	rw.RUnlock()
}

func write(n int) {
	rw.Lock()
	fmt.Printf("写goroutine %d 正在写入...\n",n)
	v := rand.Intn(1000)

	count1 = v

	fmt.Printf("写goroutine %d 写入结束，新值为：%d\n", n,v)
	wg3.Done()
	rw.Unlock()
}