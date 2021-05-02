package main

import (
	"bytes"
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	var pre bytes.Buffer
	//向bt中写入字符串
	for i := 0; i < len(strs[0]); i++ {
		tmp := strs[0][i]
		for _, s := range strs {
			if len(s) <= i || s[i] != tmp {
				return pre.String()
			}
		}
		pre.WriteByte(tmp)
	}
	return pre.String()
}

func main() {
	//strs := []string{"flower", "flow", "flight"}
	strs := []string{"ab", "a"}
	fmt.Println(strs)
	fmt.Println(longestCommonPrefix(strs))
}
//执行用时：
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗：
//2.3 MB
//, 在所有 Go 提交中击败了
//25.46%
//的用户