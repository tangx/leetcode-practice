package hascycle

// hasCycle 追缉问题
// fast 每次走2步， slow 走一步。
// 只要有环， fast 一定会追到 slow。
// 参考操场跑圈
func hasCycle(head *ListNode) bool {

	slow, fast := head, head
	for fast != nil {
		// 开始移动
		// 如果 fast 走到底了
		if fast.Next == nil {
			return false
		}

		// 移动
		slow, fast = slow.Next, fast.Next.Next

		// 如果追到了
		if slow == fast {
			return true
		}

	}
	return false
}
