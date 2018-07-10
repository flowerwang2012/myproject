package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan int, 10)
	go producer(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)
	fmt.Println("主goroutine已结束")
}

func producer(ch chan<- int){
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("send:", i)
	}
}

func consumer(ch <-chan int){
	for i := 0; i < 10; i++ {
		<- ch
		fmt.Println("receive:", i)
	}
}