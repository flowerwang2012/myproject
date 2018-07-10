package main

import (
	"fmt"
	"time"
)

func main(){
	i := 1
	go func(a int) {
		fmt.Println(a)
		fmt.Println(2)
	}(i)
	fmt.Println(3)
	time.Sleep(1 * time.Second)

	ch := make(chan int) //无缓冲channel是同步的
	go func() {
		ch <- 11
	}()
	v := <- ch //阻塞等待线程，无缓冲通道是同步的，取出就一定要等待赋值，赋值就一定要等待取出
	fmt.Println(v)

	ch2 := make(chan int)
	go func() {
		v := <- ch2
		fmt.Println(v)
	}()
	ch2 <- 22

	ch3 := make(chan int, 1) //有缓冲channel解决同步阻塞问题
	ch3 <- 33
	go func() {
		v := <- ch3
		fmt.Println(v)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("主goroutine已结束")
}