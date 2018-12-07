package main

import "fmt"
// 字符串中的第一个唯一字符
func main() {
	s := "loveleetcode"
	arr := []byte(s)
	var index int
	for i := 0; i < len(arr); i ++ {
		index = i
		repeat := false
		for j := i + 1; j < len(arr); j ++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			break
		}
	}
	fmt.Println(index)
}
