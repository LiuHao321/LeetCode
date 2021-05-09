package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
//面试题 02.01. 移除重复节点
//编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。
func removeDuplicateNodes(head *ListNode) *ListNode {
	pre, cur := head, head
	m := map[int]struct{}{}
	for cur != nil {
		if _, ok := m[cur.Val]; ok {
			pre.Next = cur.Next
		} else {
			m[cur.Val] = struct{}{}
			pre = cur
		}
		cur = cur.Next
	}
	return head
}

func main() {
	list := []int{1, 2, 3, 3, 2, 1}
	head := &ListNode{}
	cur := head
	for _, v := range list {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	removeDuplicateNodes(head)
	cur = head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}
//执行用时：
//12 ms
//, 在所有 Go 提交中击败了
//96.71%
//的用户
//内存消耗：
//6.5 MB
//, 在所有 Go 提交中击败了
//5.26%
//的用户