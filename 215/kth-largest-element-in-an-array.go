package main

import "fmt"
//数组中的第K个最大元素
//在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
type heap struct {
	Num int
	Val []int
}

func findKthLargest(nums []int, k int) int {
	h := buildMinHeap(nums[0:k])
	for i := k; i < len(nums); i++ {
		if nums[i] > h.Val[1] {
			h = add(h, nums[i])
		}
	}
	return h.Val[1]
}
func buildMinHeap(nums []int) heap {
	h := heap{Num: len(nums), Val: []int{0}}
	for _, v := range nums {
		h.Val = append(h.Val, v)
		i := len(h.Val) - 1
		for i != 1 {
			if h.Val[i/2] > h.Val[i] {
				h.Val[i/2], h.Val[i] = h.Val[i], h.Val[i/2]
			}
			i /= 2
		}
	}
	return h
}
func add(h heap, x int) heap {
	h.Val[1] = x
	i := 1
	n := h.Num
	for {
		//有双子节点
		if i*2+1 <= n {
			// 选择小的节点
			if h.Val[i*2] < h.Val[i*2+1] {
				if h.Val[i] > h.Val[i*2] {
					h.Val[i], h.Val[i*2] = h.Val[i*2], h.Val[i]
					i = i * 2
				} else {
					break
				}
			} else {
				if h.Val[i] > h.Val[i*2+1] {
					h.Val[i], h.Val[i*2+1] = h.Val[i*2+1], h.Val[i]
					i = i*2 + 1
				} else {
					break
				}
			}
		} else if i*2 == n { //有单子节点
			if h.Val[i] > h.Val[i*2] {
				h.Val[i], h.Val[i*2] = h.Val[i*2], h.Val[i]
				i = i * 2
			} else {
				break
			}
		} else {
			break
		}
	}
	return h
}
func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	fmt.Println(findKthLargest(nums, k))
}
//执行用时：
//4 ms
//, 在所有 Go 提交中击败了
//99.04%
//的用户
//内存消耗：
//3.7 MB
//, 在所有 Go 提交中击败了
//37.27%
//的用户