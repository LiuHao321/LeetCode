package main

import (
	"fmt"
	"sort"
)

// 寻找两个正序数组的中位数
// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

// 耍流氓式解法
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	if len(nums)%2 == 0 {
		return float64(nums[len(nums)/2-1]+nums[len(nums)/2]) / 2.0
	} else {
		return float64(nums[len(nums)/2])
	}
}

// 能想到的正常解法 双指针，迭代到总长度一半

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
