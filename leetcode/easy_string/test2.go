package main

import "fmt"
// 字符串中的第一个唯一字符
func main() {
	s := "loveleetcode"
	arr := []byte(s)
	var index int
	for i := 0; i < len(arr); i ++ {
		repeat := false
		for j := 0; j < len(arr); j ++ {
			if i != j {
				if arr[i] == arr[j] {
					repeat = true
					break
				}
			}
		}
		if !repeat {
			index = i
			break
		}
	}
	fmt.Println(index)
}
