package main

import "fmt"

// 无重复字符的最长子串
// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
// 两层for循环，时间和内存占用都很大，想办法优化一下
func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		ans = max(ans, rk-i+1)
	}
	return ans
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func main() {
	s := "abcabcbb"
	//s := "okojihfwuiaehf7655"
	fmt.Println(s)
	fmt.Println(lengthOfLongestSubstring(s))
}
