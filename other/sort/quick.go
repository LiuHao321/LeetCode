package main

import "fmt"

func quickSort(nums []int, start int, end int) []int {
	if start < end {
		l, r := start, end
		tag := nums[l]
		for l < r {
			for l < r && nums[l] < tag {
				l++
			}
			for l < r && nums[r] > tag{
				r--
			}
			nums[l], nums[r] = nums[r], nums[l]
		}
		quickSort(nums, start, l-1)
		quickSort(nums, l+1, end)
	}
	return nums
}
func main() {
	nums := []int{4, 2, 5, 6, 1, 9, 0}
	fmt.Println(nums)
	fmt.Println(quickSort(nums, 0, len(nums)-1))
	fmt.Println(nums)

}
//
//func quick(nums []int, start int, end int) []int  {
//	if start < end {
//		l, r := start, end
//		tag := nums[l]
//		for l < r {
//			for l < r && nums[l] < tag {
//				l++
//			}
//			for l < r && nums[r] > tag {
//				r--
//			}
//			nums[l], nums[r] = nums[r], nums[l]
//		}
//		quick(nums, start, l-1)
//		quick(nums, l+1, end)
//	}
//}