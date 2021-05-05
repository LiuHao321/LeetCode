package main

type ListNode struct {
	Val  int
	Next *ListNode
}
//合并两个有序链表
//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	cur := &ListNode{}
	res := cur
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 == nil {
		cur.Next = l2
	}
	if l2 == nil {
		cur.Next = l1
	}
	return res.Next
}
//
//执行用时：
//4 ms
//, 在所有 Go 提交中击败了
//67.23%
//的用户
//内存消耗：
//2.5 MB
//, 在所有 Go 提交中击败了
//87.03%
//的用户