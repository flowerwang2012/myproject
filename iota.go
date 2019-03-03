package main

import "fmt"

type season int

type weapon int

func main() {
	const (
		spring season = iota
		summer
		autumn
		winter
	)
	const (
		a weapon = 5 << iota // 把5左移0位
		b
		c
		d
	)
	fmt.Printf("%d %d %d %d\n", spring, summer, autumn, winter) //0 1 2 3
	fmt.Printf("%d %d %d %d\n", a, b, c, d) //5 10 20 40
	fmt.Printf("%b %b %b %b\n", a, b, c, d) //二进制101 1010 10100 101000
}
