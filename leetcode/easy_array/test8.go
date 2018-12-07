package easy_array

import "fmt"
// 移动零 不为零的都前面，这样就只剩零了
func main() {
	nums := []int{0,1,0,2,3}
	newi := 0
	for i := 0; i < len(nums); i ++ {
		if nums[i] != 0 {
			nums[newi] = nums[i]
			newi++
		}
	}
	for i := newi; i < len(nums); i ++ {
		nums[i] = 0
	}
	fmt.Println(nums)
}
