package main

type ListNode struct {
	Val int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	curr := res
	for {
		if l1 == nil {
			curr.Next = l2
			break
		} else if l2 == nil {
			curr.Next = l1
			break
		} else if l1.Val <= l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	return res.Next
}
