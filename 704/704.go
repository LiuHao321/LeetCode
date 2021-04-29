//package main
//
//import "fmt"
//
//func search(nums []int, target int) int {
//	s := 0
//	e := len(nums) - 1
//	for s <= e {
//		if nums[s+(e-s)/2] == target {
//			return s + (e-s)/2
//		}
//		if target > nums[s+(e-s)/2] {
//			s = s + (e-s)/2 + 1
//		} else {
//			e = s + (e-s)/2 - 1
//		}
//	}
//	return -1
//}
//func main() {
//	nums := []int{-1, 0, 3, 5, 9, 12}
//	target := 9
//	fmt.Println(nums)
//	fmt.Println(search(nums, target))
//}
//

package main

import "fmt"

func search(nums []int, target int) int {
	s := 0
	e := len(nums) - 1
	for s <= e {
		if nums[s+(e-s)/2] == target {
			return s + (e-s)/2
		}
		if target > nums[s+(e-s)/2] {
			s = s + (e-s)/2 + 1
		} else {
			e = s + (e-s)/2 - 1
		}
	}
	return -1
}
func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println(nums)
	fmt.Println(search(nums, target))
}
