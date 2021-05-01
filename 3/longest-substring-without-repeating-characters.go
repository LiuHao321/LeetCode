package main

import (
	"fmt"
	"strings"
)

// 无重复字符的最长子串
// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
// 核心：只增大不减小的滑动窗口
//流程：两个指针start和end表示窗口大小，遍历一次字符串，窗口在遍历过程中滑动或增大
//tips：配合画图思考更佳
//
//窗口内没有重复字符：此时判断i+1与end的关系，超过表示遍历到窗口之外了，增大窗口大小
//窗口内出现重复字符：此时两个指针都增大index+1，滑动窗口位置到重复字符(前面的字符串中)的后一位
//遍历结束，返回end-start，窗口大小
//思考：如果需要返回字符串怎么做？
//解答：只需要在窗口增大的时候记录start指针即可

// end - start 用来保存遍历过程中最大的窗口的大小，一旦超过该大小便对其进行扩展

func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += index + 1
			end += index + 1
		}
	}
	return end - start
}
func main() {
	s := "abcabcbb"
	//s := "okojihfwuiaehf7655"
	fmt.Println(s)
	fmt.Println(lengthOfLongestSubstring(s))
}

//执行用时：
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗：
//2.5 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户
