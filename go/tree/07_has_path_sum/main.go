package haspathsum

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// hasPathSum 计算是否有一条路径节点与目标值相等
// 迭代
// target-val == 0 ? return true ; return hashPathSum(subnode, target-val)
// 子层返回的 false 需要处理， 返回的 true 直接返回（最终结果）
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// 需要直接将 true 最上层结果
	// 判断该树中是否存在 **根节点到叶子节点** 的路径，这条路径上所有节点值相加等于目标和 targetSum 。
	// 叶子节点的定义是, left==nil && right==nil
	// 因此， root 节点可能是 叶子节点
	target := targetSum - root.Val
	if target == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	if hasPathSum(root.Left, target) || hasPathSum(root.Right, target) {
		return true
	}

	return false
}
