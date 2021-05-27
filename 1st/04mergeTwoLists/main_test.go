package mergetowlist04

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head = &ListNode{Val: 999, Next: nil}
	pre := head

	for {
		if l1 == nil {
			pre.Next = l2
			return head.Next
		}
		if l2 == nil {
			pre.Next = l1
			return head.Next
		}

		if l1.Val < l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
}
