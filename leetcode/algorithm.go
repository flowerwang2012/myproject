package main

import (
	"fmt"
)

// 迭代的是人，递归的是神
//–L. Peter Deutsch

// 一个数的阶乘
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 汉诺塔问题
func hanNuo(n int, a, b, c string) {
	if n == 1 { //盘子只有一个时候，从A移动到C，这也是递归的终止条件
		fmt.Printf("将盘子【%d】从 %s 移动到 %s \n", n, a, c)
	} else {
		hanNuo(n-1, a, c, b) //将a柱子上的从上到下n-1个盘移到b柱子上
		fmt.Printf("将盘子【%d】从 %s 移动到 %s \n", n, a, c)
		hanNuo(n-1, b, a, c) //将b柱子上的n-1个盘子移到c柱子上
	}
}

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
			if nums[i]+nums[j] == sum {
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

//排序算法，二分法查找
func Test3() {
	//nums := []int{2, 11, 14, 17, 20, 23, 30, 37, 40}
	nums := []int{11, 2, 23, 17, 20, 14, 30, 37, 40}
	//冒泡排序
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				tmp := nums[j]
				nums[j] = nums[i]
				nums[i] = tmp
			}
		}
	}
	fmt.Println(nums)
	//选择排序
	nums = []int{11, 2, 23, 17, 20, 14, 30, 37, 40}
	for i := 0; i < len(nums); i++ {
		m := i
		for j := i + 1; j < len(nums); j++ { //找出从i下标开始最小值的下标
			if nums[j] < nums[i] {
				m = j
			}
		}
		if i != m { //i != m说明存在需要交换位置的元素，选择元素
			tmp := nums[m]
			nums[m] = nums[i]
			nums[i] = tmp
		}
	}
	fmt.Println(nums)
	//快速排序
	nums = []int{11, 2, 23, 17, 20, 14, 30, 37, 40}
	nums = []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	nums = quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	//二分法查找
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := middleFind(nums, 3)
	//二分法查找递归
	i = middleFind2(nums, 0, len(nums)-1, 3)
	fmt.Println(i)
}
//快速排序
func quickSort(nums []int, left, right int) []int {
	value := nums[left]
	p := left //坑位，用于放置找到对应大小的值
	i, j := left, right

	for i <= j {
		for j >= p && nums[j] >= value {
			j--
		}
		if j >= p { //找到比value小的值
			nums[p] = nums[j] //将坑位放置对应的值
			p = j             //此时p代表 希望放置 比value大的值的坑位
		}

		for nums[i] <= value && i <= p {
			i++
		}
		if i <= p { //找到比value大的值
			nums[p] = nums[i] //将坑位放置对应的值
			p = i             //此时p代表 希望放置 比value小的值的坑位
		}
	}
	nums[p] = value
	if p-left > 1 {
		quickSort(nums, left, p-1)
	}
	if right-p > 1 {
		quickSort(nums, p+1, right)
	}
	return nums
}
//二分法查找
func middleFind(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
//二分法查找递归
func middleFind2(nums []int, low, high, target int) int {
	if low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			return middleFind2(nums, mid+1, high, target)
		} else {
			return middleFind2(nums, low, mid-1, target)
		}
	}
	return -1
}

//去重求合集
func Test4() {
	c1 := []rune{'a', 'b', 'c'}
	c2 := []rune{'c', 'd', 'e'}
	//如果用两层for循环迭代，时间复杂度是O(n^2)
	var new []rune
	repeat := false
	for i := 0; i < len(c1); i++ {
		for j := 0; j < len(c2); j++ {
			if c1[i] == c2[j] {
				repeat = true
				break
			}
		}
		if repeat {
			continue
		} else {
			new = append(new, c1[i])
		}
	}
	new = append(new, c2...)
	fmt.Println(string(new)) //abcde
	//所以从复杂度上，这种做法是不可取的
	//方法1：两个数组拼接成一个，通过map去重，但是map是无序的
	new = append(c1, c2...)
	m := make(map[rune]int, len(new))
	for i, v := range new {
		m[v] = i
	}
	new = []rune{}
	for k, _ := range m {
		new = append(new, k)
	}
	fmt.Println(string(new)) //deabc
	//方法2：将数组1的元素放到map中，用map来过滤掉两个数组的重复元素
	m = make(map[rune]int, len(new))
	for i, v := range c1 {
		m[v] = i
	}
	new = []rune{}
	for _, v := range c2 {
		if _, ok := m[v]; !ok {
			new = append(new, v)
		}
	}
	new = append(c1, new...)
	fmt.Println(string(new)) //abcde
}

func main() {
	nums := []int{2, 7, 11, 15}
	ret1 := Test1(nums, 9)
	ret2 := Test2(nums, 9)
	fmt.Println(ret1)
	fmt.Println(ret2)
	Test3()
	Test4()
	fmt.Println(factorial(4))
	hanNuo(3, "a", "b", "c")
}