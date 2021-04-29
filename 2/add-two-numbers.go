package main

import "fmt"

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

// 测试代码

func main() {
	num1 := []int{2, 4, 3}
	num2 := []int{5, 6, 4}
	//num1 := []int{9, 9, 9, 9, 9, 9, 9}
	//num2 := []int{9, 9, 9, 9}
	// 生成原始链表

	head1 := &ListNode{}
	cur := head1
	for _, v := range num1 {
		tmp := &ListNode{}
		tmp.Val = v
		cur.Next = tmp
		cur = cur.Next
	}

	head2 := &ListNode{}
	cur = head2
	for _, v := range num2 {
		tmp := &ListNode{}
		tmp.Val = v
		cur.Next = tmp
		cur = cur.Next
	}

	cur = head1.Next
	in1 := []int{}
	for cur != nil {
		in1 = append(in1, cur.Val)
		cur = cur.Next
	}
	fmt.Println(in1)

	cur = head2.Next
	in2 := []int{}
	for cur != nil {
		in2 = append(in2, cur.Val)
		cur = cur.Next
	}
	fmt.Println(in2)

	head3 := addTwoNumbers(head1, head2)

	// 打印处理后的链表
	cur = head3.Next
	out := []int{}
	for cur != nil {
		out = append(out, cur.Val)
		cur = cur.Next
	}
	fmt.Println(out)
}
