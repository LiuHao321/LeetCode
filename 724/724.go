package main

import "fmt"

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
