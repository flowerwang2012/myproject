package main

import (
	"fmt"
	"os"
	"os/signal"

	//"io/ioutil"
)

// 分配的内存不足以放下数据项序列,称为内存溢出
// 内存溢出，就是说，你向系统申请了装10个橘子的篮子（内存）并拿到了，但你却用它来装10个苹果，从而超出其最大能够容纳的范围，于是产生溢出；

// 内存泄漏，就是说系统的篮子（内存）是有限的，而你申请了一个篮子，拿到之后没有归还（忘记还了或是丢了），于是造成一次内存泄漏。
// 在你需要用篮子的时候，又去申请，如此反复，最终系统的篮子无法满足你的需求，最终会由内存泄漏造成内存溢出。

func main() {
	//buf, err := ioutil.ReadFile("./Go语言从入门到进阶实战.pdf")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//str := string(buf)
	//for i := 0; i < 100; i++ {
	//	fmt.Println(i)
	//	str += str + "溢出"
	//}
	sc := make(chan os.Signal, 0)
	signal.Notify(sc, os.Interrupt)

	c := make(chan int, 10)
	for i := 0; i < 10000000; i++ {
		fmt.Println("启动：", i)
		go OutOfMemory(i, c)
	}

	select {
	case <-sc:
		fmt.Println("退出程序")
		os.Exit(1)
	}
}

func OutOfMemory(i int, c chan int) {
	c <- 1
	fmt.Println(i)
}