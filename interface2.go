package main

import "fmt"

type Invoke interface {
	Call()
}
// 结构体实现接口
type Struct struct {
}
func (s Struct) Call() { //Struct结构体实现了接口，所以实例化结构体（也就是创建对象）
	fmt.Println("from struct")
}
// 函数体实现接口
// 需要将函数定义为类型，使用类型实现接口，当类型方法被调用，还需要调用函数本体
type FuncCaller func()
func (f FuncCaller) Call() {
	f()
	fmt.Println("from func")
}
func main() {
	var invoke Invoke
	invoke = new(Struct) //创建对象，invoke是接口类型的引用
	invoke.Call()
	var f FuncCaller
	f = func() {
	}
	invoke = f
	invoke.Call()
}