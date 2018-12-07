package easy_array

import "fmt"
// 两数之和
func main() {
	nums := []int{2, 7, 11, 15}
	total := 9
	isOk := false
	result := make([]int, 0)
	for i := 0; i < len(nums); i ++ {
		for j := i + 1; j < len(nums); j ++ {
			if total == nums[i] + nums[j] {
				result = append(result, nums[i], nums[j])
				isOk = true
				break
			}
		}
		if isOk {
			break
		}
	}
	fmt.Println(result)
}
