package main

import (
	"fmt"
	"time"
)

func main() {
	//发送到chan的值一旦被接收就没了，不能再次接收，所以要同时结束4个goroutine需要定义4个chan，chan通信不行，用context吧
	stop := make(chan bool)
	stop2 := make(chan bool)
	go watcher(stop, "监控1号")
	go watcher(stop2, "监控2号")
	time.Sleep(3 * time.Second)
	fmt.Println("可以了，停止所有监控")
	stop <- true
	stop2 <- true
	time.Sleep(3 * time.Second)
	fmt.Println("所有监控已经停止")
}

func watcher(stop chan bool, name string) {
	for {
		select {
		case <-stop:
			fmt.Printf("停止%s...\n", name)
			return
		default:
			fmt.Printf("%s工作中...\n", name)
			time.Sleep(1 * time.Second)
		}
	}
}