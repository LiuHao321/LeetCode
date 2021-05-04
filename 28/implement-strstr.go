package main

import (
	"fmt"
)

// 实现strStr()函数。
//
//给你两个字符串haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回 -1 。
//

func strStr(haystack string, needle string) int {
	l1 := len(haystack)
	l2 := len(needle)
	if l2 == 0 {
		return 0
	}
	if l1 == 0 || l1 < l2 {
		return -1
	}
	for i := 0; i <= l1-l2; i++ {
		if haystack[i:i+l2] == needle {
			return i
		}
	}
	return -1
}
func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println(haystack)
	fmt.Println(needle)
	fmt.Println(strStr(haystack, needle))
}

//执行用时：
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗：
//2.2 MB
//, 在所有 Go 提交中击败了
//90.43%
//的用户

//// 不要脸解法
//func strStr(haystack string, needle string) int {
//	return strings.Index(haystack, needle)
//}
//执行用时：
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗：
//2.2 MB
//, 在所有 Go 提交中击败了
//99.96%
//的用户
