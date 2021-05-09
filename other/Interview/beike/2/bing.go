package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	nums := []int{}
	for i := 0; i < n; i++ {
		x := 0
		fmt.Scan(&x)
		nums = append(nums, x)
	}
	//n := 9
	//nums := []int{2, 10, 3, 9, 2, 5, 3, 2, 9}
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	x := 0
	for _, v := range m {
		if v > x {
			x = v
		}
	}
	fmt.Println(x)
}
