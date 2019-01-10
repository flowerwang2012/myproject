package main

import "fmt"
//Go语言可以将类型的方法和普通函数视为一个概念，从而简化方法和函数混合作为回调类型时的复杂性
type Class struct {

}
//结构体方法
func (c *Class) Do(i int) {
	fmt.Println("Call method do:", i)
}
//普通函数
func funcDo(i int) {
	fmt.Println("Call function do:", i)
}

func main() {
	//声明一个回调函数
	var delegate func(int)
	//创建结构体实例
	c := new(Class)
	//设为方法
	delegate = c.Do
	delegate(100)
	//设为函数
	delegate = funcDo
	delegate(100)
}
