package main

import "fmt"

//面试题 02.06. 回文链表
//编写一个函数，检查输入的链表是否是回文的。
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 翻转后一半节点
	var rec *ListNode
	for slow != nil {
		rec = &ListNode{
			Val:  slow.Val,
			Next: rec,
		}
		slow = slow.Next
	}
	for rec != nil {
		if head.Val == rec.Val {
			head = head.Next
			rec = rec.Next
		} else {
			return false
		}
	}
	return true
}

func main() {
	//list := []int{1, 2, 3, 4, 4, 3, 2, 1}
	list := []int{1, 2}
	head := &ListNode{}
	cur := head
	for _, v := range list {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	fmt.Println(isPalindrome(head.Next))
}
//执行用时：
//12 ms
//, 在所有 Go 提交中击败了
//86.94%
//的用户
//内存消耗：
//6 MB
//, 在所有 Go 提交中击败了
//60.41%
//的用户