package main

import "fmt"

func insertSort(nums []int)  {
	for i := 1; i < len(nums); i++ {
		val := nums[i]
		index := i - 1
		for index >= 0 && nums[index] > val {
			nums[index+1] = nums[index]
			index--
		}
		if index + 1 == i {
			continue
		}
		nums[index+1] = val

	}
}

func main() {
	arr := []int{13,78,10,45,664,12}
	insertSort(arr)
	fmt.Println(arr)
}