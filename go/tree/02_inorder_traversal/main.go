package inordertraversal

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

func inorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	ret := []int{}
	// 左
	if root.Left != nil {
		ret = append(ret, inorderTraversal(root.Left)...)

	}

	// 根
	ret = append(ret, root.Val)

	// 右
	if root.Right != nil {
		ret = append(ret, inorderTraversal(root.Right)...)
	}

	return ret
}
