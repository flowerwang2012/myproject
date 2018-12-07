package main

import "fmt"
// 反转字符串
func main() {
	s := "hello"
	b := []byte(s)
	var result string
	for i := len(b) - 1; i >= 0; i -- {
		c := string(b[i])
		result = result + c
	}
	fmt.Println(result)
}
