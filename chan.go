package main

import (
	"fmt"
	"sync"
)

var (
	wg2  sync.WaitGroup
	//无缓冲的通道，发送goroutine和接收gouroutine必须是同步的，同时准备后，如果没有同时准备好的话，先执行的操作就会阻塞等待，直到另一个相对应的操作准备好为止。这种无缓冲的通道我们也称之为同步通道。
	wait = make(chan int)
)

func main() {
	wg2.Add(2)
	ch := make(chan int, 10)
	go producer(ch)
	go consumer(ch)
	//time.Sleep(1 * time.Second) 使用主goroutine休眠的方式，并不能做到等待等待其他go程运行完毕，可以使用waitgroup或者chan方式让主goroutine等待其他go程运行完毕

	//wg2.Wait()
	//fatal error: all goroutines are asleep - deadlock! 在这里阻塞等待会让接受通道<-wait无法与发送通道同时准备好，而goroutine里发送通道一直在等待接收，defer函数无法执行，从而导致死锁
	fmt.Println(<-wait)
	fmt.Println(<-wait)
	//wg2.Wait()
	fmt.Println("主goroutine已结束")
}

func producer(ch chan<- int) {
	defer wg2.Done()
	for i := 0; i < 1000000; i++ {
		ch <- i
		fmt.Println("send:", i)
	}
	wait <- 1
}

func consumer(ch <-chan int) {
	defer wg2.Done()
	for i := 0; i < 1000000; i++ {
		i = <-ch
		fmt.Println("receive:", i)
	}
	wait <- 2
}
