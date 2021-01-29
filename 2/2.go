package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0
	res := &ListNode{}
	carry := false
	tmp := res
	for true {
		if carry {
			sum = l1.Val + l2.Val + 1
		} else {
			sum = l1.Val + l2.Val
		}
		if sum >= 10 {
			tmp.Next = &ListNode{Val: sum - 10}
			carry = true
		} else {
			tmp.Next = &ListNode{Val: sum}
			carry = false
		}
		tmp = tmp.Next
		if l1.Next == nil && l2.Next == nil {
			if carry {
				tmp.Next = &ListNode{Val: 1}
			}
			break
		}

		if l1.Next == nil {
			l1.Val = 0
		} else {
			l1 = l1.Next
		}
		if l2.Next == nil {
			l2.Val = 0
		} else {
			l2 = l2.Next
		}

	}
	return res.Next
}

//func main {
//
//}
