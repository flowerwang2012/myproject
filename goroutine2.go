package main

import (
	"sync"
	"fmt"
	"time"
)
// 需求：有100个任务，但是只能起10个协程在处理
// 要求：协程处理完任务（包括失败）就去领任务继续处理直到没有任务了
// 思路：将100个任务放入channel缓存中，由于channel本身就是带锁的队列，所以多个goroutine不会产生竞争问题
var (
	tasks chan int
	runErr chan int
	wgroup sync.WaitGroup
)

func main() {
	initTask()
	wgroup.Add(10)
	for i := 0; i < 10; i++ {
		go runTask(i)
	}
	wgroup.Wait()
	// 执行结果
	if len(runErr) > 0 { // len(chan)通道里的元素数量
		for {
			select {
			case task := <-runErr:
				fmt.Printf("执行任务：%d 出错\n", task)
			default:
				//break // select里面的break是无法退出for语句的
				goto forEnd
			}
		}
		forEnd:
		fmt.Println("100个任务执行结束，有错误")
	} else {
		fmt.Println("100个任务执行结束，无错误")
	}
}
// 初始化任务
func initTask() {
	tasks = make(chan int, 100)
	runErr = make(chan int, 100)
	for i := 1; i <= 100; i++ {
		tasks <- i
	}
}
// 领取任务
func acquireTask() int {
	select {
	case i := <-tasks:
		return i
	default:
		return -1
	}
}
// 执行任务
func runTask(gNum int) {
	for {
		task := acquireTask()
		if task == -1 {
			break
		}
		// 协程需要通知哪些任务执行失败了，但不影响继续执行其他任务
		// 假设任务10, 20, 30执行出错
		if task == 10 || task == 20 || task == 30 {
			runErr <- task
			continue
		}
		fmt.Printf("开始执行任务：%d\n", task)
		time.Sleep(1 * time.Second)
		fmt.Printf("任务：%d执行完毕\n", task)
	}
	fmt.Printf("退出协程%d\n", gNum)
	wgroup.Done()
}
