package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
  Val int
  Next *ListNode
}
// 反转链表
// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
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

	o := reverseList(head)

	// 打印处理后的链表
	cur = o
	out := []int{}
	for cur.Next != nil {
		out = append(out, cur.Val)
		cur = cur.Next
	}
	fmt.Println(out)
}
