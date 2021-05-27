package chain02

/*
删除链表的倒数第N个节点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

进阶：你能尝试使用一趟扫描实现吗？



示例 1：


输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]


提示：

链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz


作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn2925/
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

func Test_RemoveNthFromEnd(t *testing.T) {
	head := &ListNode{1, nil}
	pre := head
	for i := 2; i <= 5; i++ {
		pre.Next = &ListNode{i, nil}
		pre = pre.Next
	}

	node := removeNthFromEnd(head, 1)
	for {
		if node == nil {
			break
		}

		fmt.Println(node.Val)
		if node.Next == nil {
			break
		}
		node = node.Next
	}

}
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	// 获取链表长度
	lenc := length(head)

	// 倒数第 n 个。
	/*
	   找到倒数第 n-1 个， 然后删除链表中的倒n节点。
	*/
	last := lenc - n
	// 删除头节点
	if last == 0 {
		return head.Next
	}

	// 锚定 head
	pre := head

	// 找到 倒数n-1 的节点
	for i := 0; i < last-1; i++ {
		pre = pre.Next
	}

	// 删除 倒数n 的节点
	pre.Next = pre.Next.Next
	return head
}

func length(head *ListNode) int {
	for i := 1; ; i++ {
		if head.Next == nil {
			return i
		}
		head = head.Next
	}
}
