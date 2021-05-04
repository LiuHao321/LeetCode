package main

import "fmt"

// 搜索插入位置
//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
//你可以假设数组中无重复元素。

func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := left + (right-left)/2
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(nums)
	fmt.Printf("target: %v \n", target)
	fmt.Println(searchInsert(nums, target))
}

//执行用时：
//4 ms
//, 在所有 Go 提交中击败了
//91.10%
//的用户
//内存消耗：
//3 MB
//, 在所有 Go 提交中击败了
//55.29%
//的用户
