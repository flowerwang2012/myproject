package easy_array

import "fmt"
//旋转数组 每次移动1位
func main() {
	nums := []int{1,2,3,4,5,6,7}
	revolve(nums, 2)
	fmt.Println(nums)
}

func revolve(nums []int, step int) {
	maxi := len(nums) - 1
	for step > 0 {
		last := nums[maxi]
		for i := len(nums) - 1; i >= 1; i -- {
			nums[i] = nums[i - 1]
		}
		nums[0] = last
		step --
	}
}

