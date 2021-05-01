package main

import "fmt"

// 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
// 你可以按任意顺序返回答案。

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		x := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == x {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(nums)
	fmt.Println(twoSum(nums, target))
}

// 执行用时：
//4 ms
//, 在所有 Go 提交中击败了
//90.26%
//的用户
//内存消耗：
//3 MB
//, 在所有 Go 提交中击败了
//99.99%
//的用户
