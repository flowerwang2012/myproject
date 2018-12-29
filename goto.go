package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if j == 3 {
				goto breakHere
			}
		}
	}
	fmt.Println(1) //跳过
	return //跳过
	breakHere:
		fmt.Println("goto here")
	fmt.Println("continue...")
}
