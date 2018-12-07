package main

import "fmt"
// 有效的字母异位词
func main(){
	s1 := "anagram"
	s2 := "nagaram"
	result := false
	if len(s1) != len(s2) {
		result = false
	} else {
		var b byte
		var index = -1
		for i := 0; i < len(s1); i ++ {
			if s1[i] != s2[i] && index < 0 {
				b = s1[i]
				index = i
			}
			if s1[i] != s2[i] && index >= 0 {
				if s2[i] == b && s2[index] == s1[i] {
					result = true
				}
			}
		}
	}
	fmt.Println(result)
}
