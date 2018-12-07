package easy_array

import "fmt"
//26. 删除排序数组中的重复项 因为数组已经排序了，先找不同的，按顺序放进数组里
func main() {
	nums := []int{0, 0, 1, 1, 2, 3}
	//for i := 0; i < len(nums); i ++ {
	//	for j := i + 1; j < len(nums); j ++ {
	//		if nums[i] == nums[j] {
	//			for ; j < len(nums); j ++ {
	//				nums[j-1] = nums[j]
	//			}
	//			fmt.Println(nums)
	//		}
	//	}
	//}
	removeDuplicates(nums)
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	if (len(nums) == 0) {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ { // 0 1 1 2 3
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j] // 1 2 3
		}
	}
	return i + 1
}
