package main

import "fmt"

// 最大连续 1 的个数
// 给定一个二进制数组， 计算其中最大连续 1 的个数。
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
