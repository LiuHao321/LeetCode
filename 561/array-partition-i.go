package main

import (
	"fmt"
	"sort"
)

// 数组拆分 I
//给定长度为 2n 的整数数组 nums ，你的任务是将这些数分成 n 对, 例如 (a1, b1), (a2, b2), ..., (an, bn) ，使得从 1 到 n 的 min(ai, bi) 总和最大。
//
//返回该 最大总和 。
//
//
func arrayPairSum(nums []int) (ans int) {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		ans += nums[i]
	}
	return
}
func main() {
	nums := []int{6, 2, 6, 5, 1, 2}
	fmt.Println(nums)
	fmt.Println(arrayPairSum(nums))
}
