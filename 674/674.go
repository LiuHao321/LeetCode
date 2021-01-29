package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	now := 1
	history := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			now++
		} else if now > history {
			history = now
			now = 1
		} else {
			now = 1
		}
	}
	if now > history {
		history = now
	}
	return history
}

func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println(nums)
	fmt.Println(findLengthOfLCIS(nums))
}
