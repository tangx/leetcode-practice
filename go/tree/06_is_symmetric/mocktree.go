package issymmetric

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

var (
	rootOnly   *TreeNode = initRootTree()
	symeTree   *TreeNode = initSymeTree()
	inSymeTree *TreeNode = initInSymeTree()
)

func init1_0() *TreeNode {
	tree := newNode(1)
	tree.Left = newNode(0)
	return tree
}

func initRootTree() *TreeNode {
	return newNode(1)
}

func initSymeTree() *TreeNode {
	symeTree := newNode(1)

	symeTree.Left, symeTree.Right = newNode(2), newNode(2)

	symeTree.Left.Left, symeTree.Left.Right = newNode(3), newNode(4)
	symeTree.Right.Left, symeTree.Right.Right = newNode(4), newNode(3)

	return symeTree
}

func initInSymeTree() *TreeNode {
	inSymeTree := newNode(1)

	inSymeTree.Left, inSymeTree.Right = newNode(2), newNode(2)

	inSymeTree.Left.Right = newNode(3)
	inSymeTree.Right.Right = newNode(3)

	return inSymeTree
}
