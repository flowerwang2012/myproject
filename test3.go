package main

import (
	"fmt"
	"runtime"

	//"time"
	//"time"
)

var a string
var b bool

func f1() {
	a = "abcd"
	b = true
}

func main() {
	runtime.GOMAXPROCS(1)
	go f1()
	runtime.Gosched() //main goroutine 让出cpu
	if b {
		fmt.Println(a)
	}
	//for {
	//	if b {
	//		fmt.Println(a)
	//		break
	//	}
	//}
	fmt.Println("main over")
}

