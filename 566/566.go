package main

import "fmt"

func matrixReshape(nums [][]int, r int, c int) [][]int {
	n, m := len(nums), len(nums[0])
	if r*c != n*m {
		return nums
	}
	ans := make([][]int, r)
	for i := range ans {
		ans[i] = make([]int, c)
	}
	for i := 0; i < n*m; i++ {
		ans[i/c][i%c] = nums[i/m][i%m]
	}
	return ans
}
func main() {
	nums := [][]int{{1, 2}, {3, 4}}
	fmt.Println(nums)
	fmt.Println(matrixReshape(nums, 1, 4))
}
