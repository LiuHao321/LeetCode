package main

import "fmt"

// 回文数
//给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
//回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x%10 == 0 && x != 0 {
		return false
	}
	res := 0
	for x > res {
		res = res*10 + x%10
		x /= 10
	}
	return x == res || x == res/10
}

func main() {
	x := 121
	fmt.Println(x)
	fmt.Println(isPalindrome(x))
}

//执行用时：
//12 ms
//, 在所有 Go 提交中击败了
//92.30%
//的用户
//内存消耗：
//4.9 MB
//, 在所有 Go 提交中击败了
//86.85%
//的用户
