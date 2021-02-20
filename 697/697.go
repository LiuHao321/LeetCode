package main

import "fmt"

type num struct {
	count,
	start,
	end int
}

func findShortestSubArray(nums []int) (ans int) {
	mp := map[int]num{}
	for i, v := range nums {
		if e, has := mp[v]; has {
			e.count++
			e.end = i
			mp[v] = e
		} else {
			mp[v] = num{1, i, i}
		}
	}
	maxCount := 0
	for _, e := range mp {
		if e.count > maxCount {
			maxCount, ans = e.count, e.end-e.start+1
		} else if e.count == maxCount {
			ans = min(ans, e.end-e.start+1)
		}
	}
	return
}
func min (a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	nums := []int{1, 2, 2, 3, 1}
	fmt.Println(nums)
	fmt.Println(findShortestSubArray(nums))

}
