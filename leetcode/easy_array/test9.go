package main

import "fmt"
// 两数之和是目标值，求两数的下标
func main() {
	nums := []int{2, 7, 11, 15}
	sum := 9
	//暴力法 O(n^2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == sum {
				fmt.Println([]int{i, j})
			}
		}
	}
	//利用map，一次for循环
	m := make(map[int]int)
	for i, v := range nums {
		sub := sum - v
		if j, ok := m[sub]; ok {
			fmt.Println([]int{j, i})
		} else {
			m[v] = i
		}
	}
}
