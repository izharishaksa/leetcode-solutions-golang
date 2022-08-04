package leetcode_solutions_golang

//https://leetcode.com/problems/remove-nth-node-from-end-of-list/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := getLen(head)
	pos := l - n
	if pos == 0 {
		head = head.Next
	}
	remove(head, 0, pos)
	return head
}

func getLen(head *ListNode) int {
	if head == nil {
		return 0
	}
	return getLen(head.Next) + 1
}

func remove(head *ListNode, curPos, targetPos int) {
	if head != nil {
		if curPos == targetPos-1 {
			head.Next = head.Next.Next
			remove(head.Next, curPos+1, targetPos)
			return
		}
		remove(head.Next, curPos+1, targetPos)
	}
}
