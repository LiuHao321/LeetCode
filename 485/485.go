package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	count := 0
	for _, x := range nums {
		if x == 1 {
			count++
			if count > max {
				max = count
			}
		} else {
			count = 0
		}
	}
	return max
}

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(nums)
	fmt.Println(findMaxConsecutiveOnes(nums))
}