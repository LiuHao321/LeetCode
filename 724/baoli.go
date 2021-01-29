package main

import "fmt"

func pivotIndex1(nums []int) int {
	if len(nums) < 1 {
		return -1
	}
	left := 0
	right := 0
	for i := 0; i < len(nums); i++ {
		left = 0
		right = 0
		for j := 0; j < i; j++ {
			left += nums[j]
		}
		for j := i + 1; j < len(nums); j++ {
			right += nums[j]
		}
		if left == right {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(nums)
	fmt.Println(pivotIndex1(nums))
}
