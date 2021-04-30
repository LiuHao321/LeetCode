package main

import (
	"bytes"
	"fmt"
	"strings"
)

// 重复叠加字符串匹配
//给定两个字符串 a 和 b，寻找重复叠加字符串 a 的最小次数，使得字符串 b 成为叠加后的字符串 a 的子串，如果不存在则返回 -1。
//
//注意：字符串 "abc" 重复叠加 0 次是 ""，重复叠加 1 次是 "abc"，重复叠加 2 次是 "abcabc"。
func repeatedStringMatch(a string, b string) int {
	if strings.Contains(a, b) {
		return 1
	} else {
		//定义Buffer类型
		var bt bytes.Buffer
		//向bt中写入字符串
		bt.WriteString(a)
		bt.WriteString(a)
		//获得拼接后的字符串
		aa := bt.String()
		if strings.Contains(aa, b) {
			return 2
		}
	}

	start := strings.Index(b, a)
	if start == -1 {
		return -1
	}
	l := len(a)
	end := (len(b) - start) % l
	k := (len(b) - start - end) / l
	s := 0
	e := 0
	if start != 0 {
		if start >= l {
			return -1
		}
		s = 1
		for i := 0; i < start; i++ {
			if b[i] != a[(l-start+i)%l] {
				return -1
			}
		}
	}
	if end != 0 {
		e = 1
		for i := 0; i < end; i++ {

			if b[start+k*l+i] != a[i] {
				return -1
			}
		}
	}
	var bt bytes.Buffer
	for i := 0; i < k; i++ {
		bt.WriteString(a)
		//获得拼接后的字符串
	}
	mid := bt.String()
	if mid != b[start:start+k*l] {
		return -1
	}

	return k + s + e
}

func main() {
	a := "abcd"
	b := "cdabcdab"
	//a := "abcd"
	//b := "abcdb"
	//a := "aba"
	//b := "babbbbaabbababbaaaaababbaaabbbbaaabbbababbbbabaabababaabaaabbbabababbbabababaababaaaaabbabaaaabaaaab"
	fmt.Println(repeatedStringMatch(a, b))
}
