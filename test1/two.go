package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 5, 8, 10, 50}
	tag := 4
	fmt.Println(searchInsert(nums, tag))
}

//func searchInsert(nums []int, target int) int {
//	n := len(nums)
//	left, right := 0, n-1
//	ans := n
//	for left <= right {
//		mid := (right-left)>>1 + left
//		if target <= nums[mid] {
//			ans = mid
//			right = mid - 1
//		} else {
//			left = mid + 1
//		}
//	}
//	return ans
//}

func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	res := n
	for left <= right {
		mid := (right-left)/2 + left
		if target <= nums[mid] {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}