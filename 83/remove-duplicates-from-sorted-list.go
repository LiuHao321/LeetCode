package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}


func main() {
	list := []int{1, 1, 2, 3, 3}
	head := &ListNode{}
	cur := head
	for _, v := range list {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	deleteDuplicates(head)
	cur = head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

//执行用时：
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗：
//3.1 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户