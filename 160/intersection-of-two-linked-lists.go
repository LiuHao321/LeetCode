package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}

func main() {
	listA := []int{4, 1, 8, 4, 5}
	listB := []int{5, 0, 1, 8, 4, 5}
	skipA := 2
	skipB := 3
	//listA := []int{2, 6, 4}
	//listB := []int{1, 5}
	//skipA := 3
	//skipB := 2
	headA := &ListNode{}
	headB := &ListNode{}
	curA := headA
	curB := headB
	var pos *ListNode
	for i, v := range listA {
		curA.Next = &ListNode{Val: v}
		curA = curA.Next
		if i == skipA {
			pos = curA
		}
	}
	for i, v := range listB {
		if i == skipB {
			curB.Next = pos
			break
		}
		curB.Next = &ListNode{Val: v}
		curB = curB.Next
	}
	curA = headA
	for curA != nil {
		curA = curA.Next
	}
	curB = headB
	for curB != nil {
		curB = curB.Next
	}
	p := getIntersectionNode(headA, headB)
	if p == nil {
		fmt.Println("null")
	} else {
		fmt.Printf("Reference of the node with value = %v", p.Val)
	}
}

//执行用时：
//32 ms
//, 在所有 Go 提交中击败了
//99.17%
//的用户
//内存消耗：
//7.2 MB
//, 在所有 Go 提交中击败了
//87.90%
//的用户
