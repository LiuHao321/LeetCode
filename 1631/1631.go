package main

import (
	"fmt"
)

func minimumEffortPath(heights [][]int) int {
	rows := len(heights)
	columns := len(heights[0])
	position := [][]int{{-1, 0},{0, -1},{1, 0},{0, 1}} // 左 上 右 下
	a := make([][]int, rows)
	for i := 0; i < rows; i++ {
		a[i] = make([]int, columns)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			b := make([]int, 4)
			for key, p := range position {
				x := i + p[0]
				y := j + p[1]
				n := abs(heights[i][j]-heights[x][y]) //now
				h := a[x][y]
				if n >= h {
					b[key] = n
				}else {
					b[key] = h
				}
			}
			a[i][j] = min(b)
		}
	}
	return a[rows][columns]
}
func abs(num int) int {
	if num >= 0 {
		return num
	} else {
		return -num
	}
}
func min(nums []int) int {
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		//从第二个 元素开始循环比较，如果发现有更大的，则交换
		if min > nums[i] {
			min = nums[i]
		}
	}
	return min
}

func main() {
	heights := [][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}
	fmt.Println(heights)
	fmt.Println(minimumEffortPath(heights))
}
