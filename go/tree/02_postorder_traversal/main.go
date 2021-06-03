package postordertraversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	ret := []int{}
	// 左
	if root.Left != nil {
		ret = append(ret, postorderTraversal(root.Left)...)

	}

	// 右
	if root.Right != nil {
		ret = append(ret, postorderTraversal(root.Right)...)
	}

	// 根
	ret = append(ret, root.Val)

	return ret
}
