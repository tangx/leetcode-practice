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

	return recursiveTree([]*TreeNode{root.Left, root.Right})
}

// recursiveTree 递归
// 注意: 需要在本层判断完所有所 正常/非正常 的条件
//      才能进行下一层的递归
func recursiveTree(nodes []*TreeNode) bool {
	// 根节点，直接 true
	if len(nodes) == 0 {
		return true
	}

	// 判断所有 node 是否为空
	// 如果是， 则为最后一层
	// [nil,nil,nil,nil]
	validn := 0
	for _, node := range nodes {
		if node != nil {
			validn++
		}
	}
	if validn == 0 {
		return true
	}

	// 奇数有效 node
	// if validn%2 != 0 {
	// 	return false
	// }

	n := len(nodes)
	// 因为是 从第二层开始， 就应该是偶数个节点。
	// 因为所有子节点， 空与非空都会被加入
	// 所以不再判断 n 是否为偶数， 以判断 n+1 层是否为镜像
	// if n/2 != 0 {
	// 	return false
	// }

	// 添加所有子节点
	// 记录当前 level 的所有 val
	// 虽然这里可以直到所有多少子节点， 但是不能使用数组。
	// 因为数组大小必须为 常量 [4]*TreeNode
	//           不能为 变量 [n]*TreeNode
	subnodes := []*TreeNode{}
	level := []interface{}{}
	for _, node := range nodes {
		if node == nil {
			level = append(level, nil)
			continue
		}
		level = append(level, node.Val)
		subnodes = append(subnodes, node.Left, node.Right)
	}

	// fmt.Println(level)
	for i := 0; i < n; i++ {
		if level[i] != level[n-1-i] {
			return false
		}
	}

	// 下一次递归
	// 并返回下一层的结果
	return recursiveTree(subnodes)
}
