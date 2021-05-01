package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//两数相加
//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//
//你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	sum := 0
	carry := 0
	tmp := res
	for true {
		sum = l1.Val + l2.Val + carry
		tmp.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		tmp = tmp.Next
		if l1.Next == nil && l2.Next == nil {
			if carry == 1 {
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
	//num1 := []int{2, 4, 3}
	//num2 := []int{5, 6, 4}
	num1 := []int{9, 9, 9, 9, 9, 9, 9}
	num2 := []int{9, 9, 9, 9}
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

//执行用时：
//8 ms
//, 在所有 Go 提交中击败了
//92.29%
//的用户
//内存消耗：
//4.6 MB
//, 在所有 Go 提交中击败了
//95.66%
//的用户