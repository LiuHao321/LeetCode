package main

import "fmt"

//给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。
//通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做IIII，而是IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为IX。
//
func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0
	for i := 0; i < len(s)-1; i++ {
		if m[s[i]] < m[s[i+1]] {
			res -= m[s[i]]
		} else {
			res += m[s[i]]
		}
	}
	res += m[s[len(s)-1]]
	return res
}

func main() {
	s := "LVIII"
	fmt.Println(s)
	fmt.Println(romanToInt(s))
}
//执行用时：
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗：
//3 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户