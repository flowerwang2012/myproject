package main

import "fmt"

//给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
//你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。
//示例:
//给定 nums = [2, 7, 11, 15], target = 9
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//方法1：时间复杂度 O(n^2) 空间复杂度 O(1)
func Test1(nums []int, sum int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == sum {
				return []int{i, j}
			}
		}
	}
	return nil
}
//方法2：时间和空间复杂度都是O(n)
func Test2(nums []int, sum int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		sub := sum - num
		if j, ok := m[sub]; ok {
			return []int{j, i}
		} else {
			m[num] = j
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	ret1 := Test1(nums, 9)
	ret2 := Test2(nums, 9)
	fmt.Println(ret1)
	fmt.Println(ret2)
}
