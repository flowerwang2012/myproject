package main

import (
	"testing"
	"math/rand"
)
//基准测试
//go test -bench=. -benchmem -run=none
//运行基准测试也要使用go test命令，不过我们要加上-bench=标记，它接受一个表达式作为参数，匹配基准测试的函数，.表示运行所有基准测试。
//因为默认情况下 go test 会运行单元测试，为了防止单元测试的输出影响我们查看基准测试的结果，可以使用-run=匹配一个从来没有的单元测试方法，过滤掉单元测试的输出，我们这里使用none，因为我们基本上不会创建这个名字的单元测试方法。
//测试时间默认是1秒，也就是1秒的时间内，调用了两千万次，每次调用花费117纳秒。如果想让测试运行的时间更长，可以通过-benchtime指定，比如3秒。
/*
-4，这个表示运行时对应的GOMAXPROCS的值。接着的20000000表示调用被测试代码的次数，最后的61.4 ns/op表示每次需要花费61.4纳秒。
BenchmarkTest1-4        20000000                61.4 ns/op            16 B/op          1 allocs/op
BenchmarkTest2-4         3000000               511 ns/op             318 B/op          2 allocs/op
*/
//扩大数组个数
const N = 10

func BenchmarkTest1(b *testing.B) {
	nums:=[]int{}
	for i:=0;i<N;i++{
		nums=append(nums,rand.Int())
	}
	nums=append(nums,7,2)

	b.ResetTimer()
	for i:=0;i<b.N;i++{
		Test1(nums,9)
	}
}

func BenchmarkTest2(b *testing.B) {
	nums:=[]int{}
	for i:=0;i<N;i++{
		nums=append(nums,rand.Int())
	}
	nums=append(nums,7,2)

	b.ResetTimer()
	for i:=0;i<b.N;i++{
		Test2(nums,9)
	}
}