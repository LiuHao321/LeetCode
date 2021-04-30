package main

import "fmt"

func selectSort(nums []int, k int) {
	length := len(nums)
	for i := 0; i < k; i++ {
		max := i
		for j := i+1; j < length; j++ {
			if nums[max] < nums[j] {
				max = j
			}
		}
		nums[i], nums[max] = nums[max], nums[i]
	}
	return
}
func main() {
	arr := []int{13,78,10,45,664,12}
	selectSort(arr, 2)
	fmt.Println(arr)
}
