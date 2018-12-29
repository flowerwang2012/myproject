package main

import "fmt"

func dummy(b int) int {
	var c int
	c = b
	return c
}

func void() {

}

func main() {
	var a int
	void()
	fmt.Println(a, dummy(0)) //0 0
}
//逃逸分析
// go run -gcflags "-m -l" escape_analysis.go
//gcflags参数是编译参数，其中-m表示进行内存分配分析，-l表示避免程序内联，也就是避免进行程序优化
//运行结果如下：
/*
# command-line-arguments
./escape_analysis.go:18:13: a escapes to heap
./escape_analysis.go:18:22: dummy(0) escapes to heap
./escape_analysis.go:18:13: main ... argument does not escape
0 0
*/