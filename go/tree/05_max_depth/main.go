package maxdepth

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	subDepL := maxDepth(root.Left)
	subDepR := maxDepth(root.Right)

	// 这里 return 需要 +1 是因为 sub depth
	// return max(subDepL, subDepR) + 1
	return max(subDepL, subDepR) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
