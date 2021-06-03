package levelorder

import "fmt"

// levelOrder_01 层序遍历
// 这道题的解题关键是: 如何定位 **每层** 在 nodesQ 中的起止位置。
// 才能将当前层的结果单独保存为一个 []int{}
// 因此使用双指针
// 0. nodesQ 保存所有节点
// 1. 每次节点遍历， 将当前层的结果保存到一个新的 level []int 中
// 1.2 将下一层的节点添加到 nodesQ 中
// 2. 重置新层在 nodesQ 中的起止位置
// 3. 进行下一次循环
func levelOrder_01(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	answer := [][]int{}
	nodesQ := []*TreeNode{}

	nodesQ = append(nodesQ, root)
	start, stop := 0, 1

	for start != stop {
		var level []int
		// 循环当前 level
		nodesQ, level = walk(nodesQ, start, stop)
		// 记录当前 level 的结果
		answer = append(answer, level)
		// 设定新循环起止点
		start, stop = stop, len(nodesQ)
	}

	return answer
}

// walk 传入每层的 nodes 队列切片
func walk(nodes []*TreeNode, start int, stop int) (new []*TreeNode, level []int) {
	fmt.Printf("len(nodes),start,stop => %d : %d, %d\n", len(nodes), start, stop)

	for ; start < stop; start++ {
		node := nodes[start]

		level = append(level, node.Val)

		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
	}

	return nodes, level

}
