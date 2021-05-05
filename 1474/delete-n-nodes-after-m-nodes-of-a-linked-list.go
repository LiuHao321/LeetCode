package main
type ListNode struct {
	Val int
	Next *ListNode
}
//删除链表 M 个节点之后的 N 个节点
//给定链表 head 和两个整数 m 和 n. 遍历该链表并按照如下方式删除节点:
//
//开始时以头节点作为当前节点.
//保留以当前节点开始的前 m 个节点.
//删除接下来的 n 个节点.
//重复步骤 2 和 3, 直到到达链表结尾.
//在删除了指定结点之后, 返回修改过后的链表的头节点.
func deleteNodes(head *ListNode, m int, n int) *ListNode {
	pre, cur := head, head
	for cur != nil {
		for i := 0; i < m && cur!= nil; i++ {
			pre = cur
			cur = cur.Next
		}
		for i := 0; i < n && cur!= nil; i++ {
			cur = cur.Next
		}
		pre.Next = cur
	}
	return head
}
//一次ac不写测试了
//执行用时：
//8 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗：
//6.1 MB
//, 在所有 Go 提交中击败了
//66.67%
//的用户