package main

import (
	"fmt"
)

func check(nums []int, k int) bool {
	low := 0
	high := 0
	for _, v := range nums {
		if v > k {
			high++
		} else {
			low++
		}
	}
	if high > low {
		return true
	} else {
		return false
	}
}
func main() {

	//n, k := 0, 0
	//fmt.Scan(&n, &k)
	//nums := []int{}
	//for i := 0; i < n; i++ {
	//	x := 0
	//	fmt.Scan(&x)
	//	nums = append(nums, x)
	//}
	n, k := 5, 8
	nums := []int{9, 9, 6, 0, 9}
	x := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			fmt.Printf("i: %v, j: %v\n", i, j)
			if check(nums[i:j], k) {
				if (j - i) > x {
					x = j - i
				}
				fmt.Printf("i: %v, j: %v --ok\n", i, j)

			}
		}
	}
	fmt.Println(n)
	fmt.Println(k)
	fmt.Println(nums)
	fmt.Println(x)
	//s := 3

}
