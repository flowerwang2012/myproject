package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {
	go AsyncRun()
	time.Sleep(1 * time.Second) // 因为程序会优先执行主线程，主线程执行完成后，程序会立即退出，没有多余的时间去执行子线程。如果在程序的最后让主线程休眠1秒钟，那程序就会有足够的时间去执行子线程。
	fmt.Println("主Goroutine结束")
}

func AsyncRun() {
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(j int) {
			defer wg.Done()
			fmt.Println(j)
		}(i) //每次将变量 i 的拷贝传进函数，防止闭包
	}
	wg.Wait()
}
