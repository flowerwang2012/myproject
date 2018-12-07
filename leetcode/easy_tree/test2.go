package main

import "fmt"

// 迭代的是人，递归的是神
//–L. Peter Deutsch

// 一个树的阶乘
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}
func main() {
	fmt.Println(factorial(4))
}
