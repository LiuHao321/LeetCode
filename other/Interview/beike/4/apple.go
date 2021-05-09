package main

import (
	"fmt"
)

func check(nums []int, s int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum == s {
		return true
	}
	for len(nums) != 0{
		avg := sum / len(nums)
		for _, v := range nums {
			if v > avg {
				nums.
			}
		}
	}

	return true
}
func main() {

	n, q := 0, 0
	fmt.Scan(&n, &q)
	nums := []int{}
	for i := 0; i < n; i++ {
		x := 0
		fmt.Scan(&x)
		nums = append(nums, x)
	}
	n, q := 5, 3
	nums := []int{7, 2, 1, 6, 5}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	avg := sum / len(nums)
	fmt.Println(n)
	fmt.Println(q)
	fmt.Println(nums)
	fmt.Println(avg)
	//s := 3

}
