package main

import "fmt"

// 二分查找
//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
func search(nums []int, target int) int {
	s := 0
	e := len(nums) - 1
	for s <= e {
		if nums[s+(e-s)/2] == target {
			return s + (e-s)/2
		}
		if target > nums[s+(e-s)/2] {
			s = s + (e-s)/2 + 1
		} else {
			e = s + (e-s)/2 - 1
		}
	}
	return -1
}
func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println(nums)
	fmt.Println(search(nums, target))
}
