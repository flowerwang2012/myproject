package main

import (
	"strconv"
	"fmt"
)
// 加一
func main() {
	num := 1234
	nums := convertNum(num)
	fmt.Println(nums)
}

func convertNum(num int) []int {
	num++
	str := strconv.Itoa(num)
	nums := make([]int, len(str))
	for i := 0; i < len(str); i ++ {
		elem := string([]byte(str)[i:i+1])
		nums[i], _ = strconv.Atoi(elem)
	}
	return nums
}
