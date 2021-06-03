package levelorder

// levelOrder 层序遍历
// 不再使用全局数组 Q
// 每次遍历， 都返回一个全新的下层数组
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	// 初始化第一层
	answers := [][]int{}
	nodes := []*TreeNode{root}

	// 每层循环
	// 直到下一层没有节点
	for {
		var answer []int
		nodes, answer = walk2(nodes)
		if len(answer) == 0 {
			break
		}
		answers = append(answers, answer)
	}

	return answers
}

// walk2 遍历 level 层次
// 遍历当前节点， 返回子节点 与 当前节点的值
func walk2(nodes []*TreeNode) (subnodes []*TreeNode, answer []int) {
	if nodes == nil {
		return
	}

	for _, node := range nodes {
		answer = append(answer, node.Val)
		if node.Left != nil {
			subnodes = append(subnodes, node.Left)
		}
		if node.Right != nil {
			subnodes = append(subnodes, node.Right)
		}
	}

	return subnodes, answer
}
