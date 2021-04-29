package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 删除链表的倒数第 N 个结点
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
// 题目要点：双指针 由于我们需要找到倒数第n个节点，因此我们可以使用两个指针first和second同时对链表进行遍历，并且first比second超前 nn 个节点。当first遍历到链表的末尾时，second就恰好处于倒数第n个节点。

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
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

	removeNthFromEnd(head, 3)

	// 打印处理后的链表
	cur = head.Next
	out := []int{}
	for cur != nil {
		out = append(out, cur.Val)
		cur = cur.Next
	}
	fmt.Println(out)
}
