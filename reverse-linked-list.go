package leetcode_solutions_golang

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.com/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return cur
}
