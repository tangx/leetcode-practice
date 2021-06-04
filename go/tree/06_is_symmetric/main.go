package issymmetric

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// isSymmetric 对称树
// 对称树的每一层都是对称的。
// 如果当前层对称则进入下一层，否则直接返回 true
// 注意: 如果最底层没有元素，返回 true
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return true
}

// 递归
func recursiveTree(nodes []*TreeNode) bool {
	// 根节点
	if len(nodes) == 1 {
		return true
	}

	// 因为是 从第二层开始， 就应该是偶数个节点。
	if len(nodes)/2 != 0 {
		return false
	}

	level := []interface{}{}

	for _, node := range nodes {

	}

}
