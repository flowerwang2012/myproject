package main

import (
	"fmt"
)
// 两个数组的交集 II
func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	result := []int{}
	for i := 0; i < len(nums1); i ++ {
		for j := 0; j < len(nums2); j ++ {
			if nums1[i] == nums2[j] {
				result = append(result, nums1[i])
				break
			}
		}
	}
	fmt.Println(result)
}
