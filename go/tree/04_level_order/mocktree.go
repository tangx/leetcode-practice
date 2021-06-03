package levelorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var MockNode *TreeNode

func newNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func init() {
	MockNode = newNode(3)

	node9 := newNode(9)
	node20 := newNode(20)
	MockNode.Left = node9
	MockNode.Right = node20

	node20.Left = newNode(15)
	node20.Right = newNode(7)
}
