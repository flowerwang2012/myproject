package main

import "fmt"
// 只出现一次的数字 每个数都要跟除自身之外的数相比较
func main() {
	nums := []int{1,2,1,4,2}
	var num int
	for i := 0; i < len(nums); i ++ {
		norepeat := true
		for j := 0; j < len(nums); j ++ {
			if i != j {
				if nums[i] == nums[j] {
					norepeat = false
					break
				}
			}
		}
		if norepeat {
			num = nums[i]
			break
		}
	}
	fmt.Println(num)
}
