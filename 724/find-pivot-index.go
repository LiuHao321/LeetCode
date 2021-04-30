package main

import "fmt"
// 寻找数组的中心下标
//给你一个整数数组 nums，请编写一个能够返回数组 “中心下标” 的方法。
//
//数组 中心下标 是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。
//
//如果数组不存在中心下标，返回 -1 。如果数组有多个中心下标，应该返回最靠近左边的那一个。
//
//注意：中心下标可能出现在数组的两端。
func pivotIndex(nums []int) int {
	if len(nums) < 1 {
		return -1
	}
	left := 0
	right := 0
	for i := 1; i < len(nums); i++ {
		right += nums[i]
	}
	if left == right {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		left += nums[i - 1]
		right -= nums[i]
		if left == right {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(nums)
	fmt.Println(pivotIndex(nums))
}
