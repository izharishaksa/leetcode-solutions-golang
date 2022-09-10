//https://leetcode.com/problems/add-two-numbers/

package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var nodes []ListNode
	res := 0
	for l1 != nil || l2 != nil {
		nodes = append(nodes, ListNode{Val: res})
		res = 0
		index := len(nodes) - 1
		if l1 != nil {
			nodes[index].Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			nodes[index].Val += l2.Val
			l2 = l2.Next
		}
		if nodes[index].Val > 9 {
			val := nodes[index].Val
			nodes[index].Val = val % 10
			res = val / 10
		}
	}
	if res > 0 {
		nodes = append(nodes, ListNode{Val: res})
	}
	for i := 1; i < len(nodes); i++ {
		nodes[i-1].Next = &nodes[i]
	}
	return &nodes[0]
}
