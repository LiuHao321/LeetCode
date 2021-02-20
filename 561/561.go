package main

import (
	"fmt"
	"sort"
)

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
