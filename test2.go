package main

import "fmt"

func main() {
	var (
		x = 91
		y = 100
		count = 0
	)
	for y > 0 {
		if x > 100 {
			x = x - 10
			y--
			fmt.Println(x,y,count)
		} else {
			x++
		}
		count++
	}
}