package main
type ListNode struct {
	Val int
	Next *ListNode
}
//链表的中间结点
//给定一个头结点为 head 的非空单链表，返回链表的中间结点。
//
//如果有两个中间结点，则返回第二个中间结点。
func middleNode(head *ListNode) *ListNode {
	first := head
	second := head
	for second != nil && second.Next!= nil {
		second = second.Next.Next
		first = first.Next
	}
	return first
}
//这题一次ac，不想本地测试了
//执行用时：
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗：
//1.9 MB
//, 在所有 Go 提交中击败了
//59.45%
//的用户