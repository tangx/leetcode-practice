package reverselist03

/*
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

相关标签
链表

Go



作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnnhm6/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

import (
	"fmt"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(head *ListNode, n int) *ListNode {
	new := &ListNode{Val: n, Next: nil}
	if head == nil {
		return new
	}

	head.Next = new
	return new
}

func Test_RemoveNthFromEnd(t *testing.T) {
	// head := &ListNode{1, nil}
	// pre := head
	// for i := 2; i <= 5; i++ {
	// 	pre.Next = &ListNode{i, nil}
	// 	pre = pre.Next
	// }
	head := NewListNode(nil, 1)
	next := NewListNode(head, 2)
	next = NewListNode(next, 3)
	next = NewListNode(next, 4)
	NewListNode(next, 5)

	pre := head
	for {
		fmt.Println("pre.Val=", pre.Val)
		if pre.Next == nil {
			break
		}
		pre = pre.Next
	}

	// ret := reverseList(head)
	ret := reverseList2(head)

	for {
		fmt.Println("ret.Val=", ret.Val)
		if ret.Next == nil {
			break
		}
		ret = ret.Next
	}

}

func reverseList2(head *ListNode) *ListNode {
	// digui
	if head == nil || head.Next == nil {
		return head
	}
	ret := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}

// 反转 链表
func reverseList(head *ListNode) *ListNode {

	// 空
	if head == nil {
		return nil
	}

	// 迭代
	nodeSet := []*ListNode{}

	pre := head
	for {
		nodeSet = append(nodeSet, pre)
		if pre.Next == nil {
			break
		}
		pre = pre.Next
	}

	tail := nodeSet[len(nodeSet)-1]
	last := tail
	for n := len(nodeSet); n >= 0; n-- {
		if n == 0 {
			last.Next = nil
			continue
		}

		last.Next = nodeSet[n-1]
		last = nodeSet[n-1]

	}

	return tail

}
