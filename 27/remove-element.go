package main

import "fmt"

// 移除元素
//给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
//
//不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
//
//元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

// 要点：双指针

func removeElement(nums []int, val int) int {
	right := len(nums) - 1
	for i := 0; i <= right; i++ {
		for {
			if right < i {
				return right + 1
			}
			if nums[right] == val {
				right--
			} else {
				break
			}
		}
		if nums[i] == val {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		}
	}
	return right + 1
}

func main() {
	//nums := []int{3, 2, 2, 3}
	//val := 3
	nums := []int{2, 2, 3}
	val := 2
	fmt.Println(nums)
	x := removeElement(nums, val)
	fmt.Println(x)
	fmt.Println(nums[0:x])
}

//执行用时：
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗：
//2.1 MB
//, 在所有 Go 提交中击败了
//88.05%
//的用户
