package haspathsum

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(n interface{}) *TreeNode {
	if i, ok := n.(int); ok {
		return &TreeNode{Val: i}
	}
	return nil
}

func generator(l []interface{}) (*TreeNode, error) {
	if len(l) == 0 {
		return nil, fmt.Errorf("nil list")
	}

	if l[0] == nil {
		return nil, fmt.Errorf("nil root node")
	}

	// 创建所有节点
	nodes := []*TreeNode{}
	for i := 0; i < len(l); i++ {
		nodes = append(nodes, NewTreeNode(l[i]))
	}

	// 顺序关联所有节点
	for cur, fast := 0, 1; fast < len(l); {
		node := nodes[cur]
		for ; node == nil; cur++ {
			if cur >= fast {
				return nodes[0], nil
			}
			node = nodes[cur]
		}

		node.Left = nodes[fast]
		if fast+1 < len(l) {
			node.Right = nodes[fast+1]
		}
		cur, fast = cur+1, fast+2

	}

	return nodes[0], nil
}
