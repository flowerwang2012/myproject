package main

import "fmt"

/*
	函数 + 引用环境（引用变量）= 闭包，多用于工厂模式，具有封装性
*/
func Accumulator(value int) func() int {
	return func() int { //闭包函数
		value++
		return value
	}
}

func main() {
	accumulator := Accumulator(1)
	fmt.Println(accumulator()) //变量会跟随闭包生命期一直存在，闭包本身就如同变量一样拥有了记忆效应
	fmt.Println(accumulator())
	fmt.Printf("%p\n", accumulator)
	accumulator2 := Accumulator(10)
	fmt.Println(accumulator2())
	fmt.Printf("%p\n", accumulator2)
}

