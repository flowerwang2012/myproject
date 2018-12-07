package easy_array

import "fmt"

//买卖股票的最佳时机 使用贪心算法，只要可以产生利润（后一天比前一天股票价值上升），就进行一次买卖
func main() {
	var lr int
	nums := []int{7,1,5,3,6,4}
	for i := 0; i < len(nums) - 1; i ++ {
		fmt.Println(nums[i],nums[i+1])
		if nums[i+1] > nums[i] {
			lr += nums[i+1] - nums[i]
		}
	}
	fmt.Println(lr)
}
