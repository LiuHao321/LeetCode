package main

type ListNode struct {
	Val  int
	Next *ListNode
}
//二进制链表转整数
//给你一个单链表的引用结点 head。链表中每个结点的值不是 0 就是 1。已知此链表是一个整数数字的二进制表示形式。
//
//请你返回该链表所表示数字的 十进制值 。
func getDecimalValue(head *ListNode) int {
	res := 0
	for head != nil {
		res = res*2 + head.Val
		head = head.Next
	}
	return res
}
//一次ac不写测试了

//执行用时：
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗：
//2 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户
