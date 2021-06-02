package preordertraversal

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

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ret := []int{}
	ret = append(ret, root.Val)

	if root.Left != nil {
		ret = append(ret, preorderTraversal(root.Left)...)

	}

	if root.Right != nil {
		ret = append(ret, preorderTraversal(root.Right)...)
	}

	return ret
}
