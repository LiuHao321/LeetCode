package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
// 删除链表中的节点
// 请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点。传入函数的唯一参数为 要被删除的节点 。
// 题目要点：给出要删除的节点，无法使用传统的从上一个节点链接到下一个节点的方法，改为将下一个节点赋值给本节点，删除下一个节点
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
// 测试代码
func main() {
	nums := []int{1, 2, 3, 4, 5}
	// 生成原始链表
	head := &ListNode{}
	cur := head
	for _, v := range nums {
		tmp := &ListNode{}
		tmp.Val = v
		cur.Next = tmp
		cur = cur.Next
	}
	cur = head.Next
	in := []int{}
	for cur != nil {
		in = append(in, cur.Val)
		cur = cur.Next
	}
	fmt.Println(in)

	cur = head
	for i := 0; i < 3; i++ {
		cur = cur.Next
	}
	deleteNode(cur)

	// 打印处理后的链表
	cur = head.Next
	out := []int{}
	for cur != nil {
		out = append(out, cur.Val)
		cur = cur.Next
	}
	fmt.Println(out)

}
