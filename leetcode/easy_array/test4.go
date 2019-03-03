package main

import "fmt"
// 存在重复
func main() {
	nums := []int{1,2,3,4}
	repeat := false
	for i := 0; i < len(nums); i ++ {
		for j := i + 1; j < len(nums); j ++ {
			if nums[i] == nums[j] {
				repeat = true
			}
		}
	}
	fmt.Println(repeat)
}
