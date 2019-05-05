package main

import (
	"sync"
	"runtime"
	"fmt"
	"sync/atomic"
)

var (
	count int32
	wg1   sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	wg1.Add(2)
	go incCount()
	go incCount()
	wg1.Wait()
	fmt.Println(count)
}

//runtime.Gosched()是让当前goroutine暂停的意思，退回执行队列，让其他等待的goroutine运行，目的是让我们演示资源竞争的结果更明显。注意，这里还会牵涉到CPU问题，多核会并行，那么资源竞争的效果更明显。
func incCount() {
	defer wg1.Done()
	for i := 0; i < 2; i++ {
		val := count
		runtime.Gosched()
		val++
		count = val
	}
}

// 留意这里atomic.LoadInt32和atomic.StoreInt32两个函数，一个读取int32类型变量的值，一个是修改int32类型变量的值，这两个都是原子性的操作
// Go已经帮助我们在底层使用加锁机制，保证了共享资源的同步和安全，所以我们可以得到正确的结果
func incCount2() {
	defer wg1.Done()
	for i := 0; i < 2; i++ {
		val := atomic.LoadInt32(&count)
		runtime.Gosched()
		val++
		atomic.StoreInt32(&count, val)
	}
}

//新声明了一个互斥锁mutex sync.Mutex，这个互斥锁有两个方法，一个是mutex.Lock(),一个是mutex.Unlock(),这两个之间的区域就是临界区，临界区的代码是安全的。
func incCount3() {
	defer wg1.Done()
	for i := 0; i < 2; i++ {
		mutex.Lock()
		val := count
		runtime.Gosched()
		val++
		count = val
		mutex.Unlock()
	}
}
