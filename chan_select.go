package main

import (
	"fmt"
	"time"
)

func main() {
	//发送到chan的值一旦被接收就没了，不能再次接收，所以要同时结束4个goroutine需要定义4个chan，chan通信不行，用context吧
	//stop := make(chan bool)
	//stop2 := make(chan bool)
	//go watcher(stop, "监控1号")
	//go watcher(stop2, "监控2号")
	//time.Sleep(3 * time.Second)
	//fmt.Println("可以了，停止所有监控")
	//stop <- true
	//stop2 <- true
	//time.Sleep(3 * time.Second)
	//fmt.Println("所有监控已经停止")
	syncChannel()
	time.Sleep(3 * time.Second)
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

func syncChannel() {
	c := make(chan int)
	go selectCase(c)
	c <- 1
	fmt.Println("开始写入数据1")
	c <- 2
	fmt.Println("开始写入数据2")
	c <- 3
	fmt.Println("开始写入数据3")
	time.Sleep(3*time.Second) //等待协程打印
	fmt.Println("三个数据读取完毕")
	c <- -1
}
func selectCase(c chan int) {
	//forEnd: //break forEnd 执行到这里
	for {	// 没有for，select只会执行一次，但也不能让这个协程一直for循环下去，需要退出for语句，结束协程
		select {
		case i := <-c: // case里的代码块，没有执行完毕，不能往通道写数据
			if i == -1 {
				//break // select里面的break是无法退出for语句的，可以使用goto语句/break标签/return语句
				goto forEnd
				//break forEnd
				//return
			}
			time.Sleep(2 * time.Second)
			fmt.Println("数据",i,"处理完毕")
		}
	}
	forEnd: //goto forEnd 执行到这里
	fmt.Println("协程selectCase结束")
}