package main

import "fmt"

//Go 结构体实现了接口方法就是实现了接口，这与Java不同，Java是先定义接口实现类，再定义接口方法的实现
//Go 一个接口可以有多个实现，一个结构体可以实现多个接口，这与Java相同，实现了多态
//Go 结构体实现了接口，创建的对象可以是结构体指针类型的引用，也可以是接口类型的引用，不同类型的引用，所能调用的方法是不同的，这与Java的唯一区别是，Java创建对象后，类型是类，Go是指针

type Human interface {
	SayHi()
}

type Person interface {
	Human
	Sing(lyrics string)
}

type Student struct {
	name string
	score float64
}

func (s Student) SayHi() { //Student结构体实现了接口，所以实例化结构体（也就是创建对象）
	fmt.Printf("student [%s, %f] say hi\n", s.name, s.score)
}

func (s Student) Sing(lyrics string) {
	fmt.Printf("student [%s, %f] sing [%s]\n", s.name, s.score, lyrics)
}

func main() {
	s := &Student{"sam", 99.99} //字面量创建 或 new()创建对象，s是对象的引用
	var p Person
	p = s
	p.SayHi()
	p.Sing("********")
	var h Human //Person为超集，Human为子集
	h = s
	h.SayHi()
	h = p
	h.SayHi()
}