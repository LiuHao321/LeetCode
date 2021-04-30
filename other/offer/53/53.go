package main

import "fmt"

func search(nums []int) []int {
	mp := map[int]int{}
	for _, v := range nums {
		mp[v]++
	}
	max := nums[0]
	max_count := mp[nums[0]]
	for i, v := range mp {
		if v > max_count {
			max = i
			max_count = v
		}
	}
	return []int{max, max_count}
}
func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(nums)
	fmt.Println(search(nums))
}
//这题没按提要求写，练了一下map的用法