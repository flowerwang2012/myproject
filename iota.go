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
		a weapon = 1 << iota // 把1左移0位
		b
		c
		d
	)
	fmt.Printf("%d %d %d %d\n", spring, summer, autumn, winter) //0 1 2 3
	fmt.Printf("%d %d %d %d\n", a, b, c, d) //1 2 4 8
	fmt.Printf("%b %b %b %b\n", a, b, c, d) //二进制1 10 100 1000，2的0次方...
}
