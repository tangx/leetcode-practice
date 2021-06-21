package ispalindromechain

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {

	// 双指针找到中心点
	// 循环完成后， slow 则在中心点
	// 1 2 3 4 5
	//     s   f
	// 1 2 3 4
	//   s f
	slow, fast := head, head
	for {
		if fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
			continue
		}
		break
	}

	// 指针重置
	// 1 2 3 4 5
	// s     f
	// 1 2 3 4
	// s   f
	slow, fast = head, slow

	// 翻转后半部分链表
	// 1 2 3 5 4
	// s     f
	// 1 2 4 3
	// s   f
	fast = reverse(fast)
	// 1 2 4 3
	//   s   f
	// 如果不到最后一个节点, 一直往后移动
	for fast != nil && slow != nil {
		// 如果相等， 向后移动
		if slow.Val == fast.Val {
			slow, fast = slow.Next, fast.Next
			continue
		}
		// 如果不相等， 返回 false
		return false
	}

	return true
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	ret := reverse(head.Next)

	head.Next.Next = head
	head.Next = nil

	return ret
}
