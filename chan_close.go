package main

import (
	"fmt"
	"time"
)

var (
	ch1 = make(chan int)
	ch2 = make(chan int, 10)
)
// 无缓冲 关闭通道 结果：直接抛出异常
// 有缓冲 关闭通道 结果：不抛异常，可以将通道中的缓冲数据读完
func main() {
	//go sendCh1() //进入等待发送队列，sendCh1阻塞
	//time.Sleep(2*time.Second)
	//close(ch1) //关闭通道
	//go recvCh1() //获取等待发送队列中的sendCh1，sendCh1开始发送元素，抛出异常
	//time.Sleep(2*time.Second)
	go sendCh2()
	time.Sleep(2*time.Second)
	close(ch2)
	go recvCh2()
	time.Sleep(2*time.Second)
	fmt.Println("程序结束")
}

func sendCh1() {
	ch1 <- 1 //send on closed channel
}

func recvCh1() {
	for {
		time.Sleep(200*time.Millisecond)
		select {
		case r, ok := <-ch1: //如果通道没有关闭，那第二个参数为true，否则为false
			if !ok {
				fmt.Println("ch1 已关闭")
			} else {
				fmt.Println("recv ch1:", r)
			}
		}
	}
}

func sendCh2() {
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
}

func recvCh2() {
	for {
		time.Sleep(200*time.Millisecond)
		select {
		case r, ok := <-ch2: //如果通道没有关闭，那第二个参数为true，否则为false
			if !ok {
				fmt.Println("ch2 已关闭")
			} else {
				fmt.Println("recv ch2:", r)
			}
		}
	}
}