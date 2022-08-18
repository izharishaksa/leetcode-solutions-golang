//https://leetcode.com/problems/merge-two-sorted-lists/

package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	cur := result
	for {
		if list1 == nil || list2 == nil {
			break
		}
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else if list1.Val > list2.Val {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return result.Next
}
