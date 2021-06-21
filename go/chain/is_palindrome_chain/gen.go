package ispalindromechain

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(i int) *ListNode {
	return &ListNode{
		Val: i,
	}
}

func NewChain(list []int) *ListNode {
	head := &ListNode{}

	ptr := head
	for _, i := range list {
		// fmt.Println(i)
		node := NewListNode(i)
		ptr.Next = node
		ptr = ptr.Next
	}

	return head.Next
}
