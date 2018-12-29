package main

import "fmt"

type Data struct {
	complax  []int // 测试切片在参数传递中的效果
	instance InnerData // 实例分配的InnerData
	ptr      *InnerData // 将ptr分配为InnerData的指针类型
}
type InnerData struct {
	a int
}
func passByValue(inFunc Data) Data {
	fmt.Printf("inFunc value: %+v\n", inFunc) // 输出变量的详细结构
	fmt.Printf("inFunc ptr: %p\n", &inFunc) // 参数inFunc的指针地址，在计算机中，拥有相同地址且类型相同的变量，表示同一块内存区域
	return inFunc //返回的过程将发生值复制
}
func main() {
	var d = Data{
		complax:[]int{1,2,3},
		instance:InnerData{
			a:100,
		},
		ptr:&InnerData{200},
	}
	fmt.Printf("in value: %+v\n", d)
	fmt.Printf("in ptr: %p\n", &d)
	var f func(inFunc Data) Data
	f = passByValue
	out := f(d)
	fmt.Printf("out value: %+v\n", out)
	fmt.Printf("out ptr: %p\n", &out)

}
//in value: {complax:[1 2 3] instance:{a:100} ptr:0xc420084008}
//in ptr: 0xc42007e180
//inFunc value: {complax:[1 2 3] instance:{a:100} ptr:0xc420084008}
//inFunc ptr: 0xc42007e210
//out value: {complax:[1 2 3] instance:{a:100} ptr:0xc420084008}
//out ptr: 0xc42007e1e0
//所有Data结构的地址发生了变化，意味所有结构都是一块新的内存
//所有Data结构的成员值都没有变化，参数是值传递
//Data结构的ptr成员在传递过程中保持一致，表示指针在函数参数值传递中传递的只是指针值，不会复制指针指向的部分