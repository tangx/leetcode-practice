package hascycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(i int) *ListNode {
	return &ListNode{
		Val: i,
	}
}

func NewChain(list []int, pos int) *ListNode {
	head := &ListNode{}

	ptr := head

	// 环的起点
	var cycleStPos *ListNode

	for key, val := range list {
		// fmt.Println(i)
		node := NewListNode(val)
		if key == pos {
			cycleStPos = node
		}
		ptr.Next = node
		ptr = ptr.Next
	}

	// 成环
	if pos != -1 {
		ptr.Next = cycleStPos
	}

	return head.Next
}
