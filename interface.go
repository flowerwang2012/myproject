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

func (s *Student) SayHi() {
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