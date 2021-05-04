package main

import "fmt"

func countAndSay(n int) string {
	res := []byte{'1'}
	for i := 1; i < n; i++ {
		tmp := []byte{}
		for j := 0; j < len(res); j++ {
			count := 1
			for j < len(res)-1 && res[j] == res[j+1] {
				count++
				j++
			}
			tmp = append(tmp, byte(count+'0'), res[j])
		}
		res = tmp
	}
	return string(res)
}

func main() {
	fmt.Println(countAndSay(5))
}

//执行用时：
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗：
//2.4 MB
//, 在所有 Go 提交中击败了
//76.29%
//的用户

// func countAndSay(n int) string {
//	res := []byte{'1'}
//	for i := 1; i < n; i++ {
//		tmp := []byte{}
//		for j := 0; j < len(res); j++ {
//			count := 1
//			for j < len(res)-1 && res[j] == res[j+1] {
//				count++
//				j++
//			}
//			tmp = append(tmp, byte(count+'0'), res[j])
//		}
//		res = tmp
//	}
//	return string(res)
//}
