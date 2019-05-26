package main

import (
	"fmt"
	"encoding/hex"
)

func main() {
	//var (
	//	x = 91
	//	y = 100
	//	count = 0
	//)
	//for y > 0 {
	//	if x > 100 {
	//		x = x - 10
	//		y--
	//		fmt.Println(x,y,count)
	//	} else {
	//		x++
	//	}
	//	count++
	//}

	s := "hello"
	i := HashCode(s)
	fmt.Println(i)
	hex.EncodeToString([]byte(s))
	nums := []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	fmt.Println(quickSort(nums, 0, len(nums)-1))
}

//进制转换+常量+位运算
func HashCode(key string) int {
	var length = 16
	var index = 0
	for k := 0; k < len(key); k++ {
		fmt.Println(int(key[k]))
		index *= (1103515245 + int(key[k])) //常量+十进制
	}
	index >>= 27
	index &= length - 1
	return index
}

func quickSort(nums []int, left, right int) []int {
	 value := nums[left]
	 p := left
	 l, r := left, right
	 for l != r {
	 	for l < r && nums[r] >= value {
	 		r--
		}
		if nums[r] < value {
			nums[p] = nums[r]
			p = r
		}
		for l < r && nums[l] <= value {
			l++
		}
		if nums[l] > value {
			nums[p] = nums[l]
			p = l
		}
	 }
	 nums[p] = value
	 if p - left > 1 {
	 	quickSort(nums, left, p-1)
	 }
	 if right - p > 1 {
		 quickSort(nums, p + 1, right)
	 }
	 return nums
}