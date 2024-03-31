package add_two_numbers

import . "leetcode-go/util"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := &ListNode{}
	curr := dummy
	carry := 0
	for l1 != nil && l2 != nil {
		curr.Next = &ListNode{Val: (l1.Val + l2.Val + carry) % 10}
		curr = curr.Next
		carry = (l1.Val + l2.Val + carry) / 10
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		curr.Next = &ListNode{Val: (l1.Val + carry) % 10}
		carry = (l1.Val + carry) / 10
		curr = curr.Next
		l1 = l1.Next
	}

	for l2 != nil {
		curr.Next = &ListNode{Val: (l2.Val + carry) % 10}
		carry = (l2.Val + carry) / 10
		curr = curr.Next
		l2 = l2.Next
	}
	if carry != 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}
