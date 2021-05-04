package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return head
	}
	pre := head
	cur := head
	for cur.Next != nil {
		if cur.Val == val {
			cur.Val = cur.Next.Val
			fmt.Println(cur.Next.Next)
			cur.Next = cur.Next.Next
		} else {
			pre = cur

			cur = cur.Next
		}
	}
	if cur.Val == val {
		pre.Next = nil
	}
	return head
}
func main() {
	//list := []int{1, 2, 6, 3, 4, 5, 6}
	//val := 6
	list := []int{5, 4, 3, 2, 1, 1}
	val := 1
	head := &ListNode{}
	cur := head
	for _, v := range list {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	removeElements(head, val)
	cur = head
	for cur != nil {
		fmt.Println(cur.Val)
		fmt.Println(cur)
		cur = cur.Next
	}
}

//执行用时：
//4 ms
//, 在所有 Go 提交中击败了
//99.62%
//的用户
//内存消耗：
//4.6 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户