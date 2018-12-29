package main

import "fmt"

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

func (s *Student) SayHi() { //这里是指针类型实现了接口，所以初始化结构体取指针
	fmt.Printf("student [%s, %f] say hi\n", s.name, s.score)
}

func (s *Student) Sing(lyrics string) {
	fmt.Printf("student [%s, %f] sing [%s]\n", s.name, s.score, lyrics)
}

func main() {
	s := &Student{"sam", 99.99}
	var p Person
	p = s
	p.SayHi()
	p.Sing("********")
	var h Human //Person为超集，Human为子集
	h = s //多态
	h.SayHi()
	h = p
	h.SayHi()
}