/*
路径总和
给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。

叶子节点 是指没有子节点的节点。



示例 1：


输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
输出：true
示例 2：


输入：root = [1,2,3], targetSum = 5
输出：false
示例 3：

输入：root = [1,2], targetSum = 0
输出：false


提示：

树中节点的数目在范围 [0, 5000] 内
-1000 <= Node.val <= 1000
-1000 <= targetSum <= 1000
相关标签
树
深度优先搜索

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/data-structure-binary-tree/xo566j/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
package haspathsum

import (
	"testing"

	. "github.com/onsi/gomega"
)

var MockDatas = []struct {
	s      string
	node   []interface{}
	target int
	valid  bool
}{
	{
		s: "1,2,4, want 3", node: []interface{}{1, 2, 4, nil},
		target: 3, valid: true,
	},
	{
		s: "true node", node: []interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1},
		target: 22, valid: true,
	},
	{
		s: "false node", node: []interface{}{1, 2, 3},
		target: 5, valid: false,
	},
	{
		s: "根节点到叶子节点 的路径, 根不能直接为结果", node: []interface{}{1, 2},
		target: 1, valid: false,
	},
	{
		s: "只有 root 节点, root 节点也是叶子节点", node: []interface{}{1},
		target: 1, valid: true,
	},
	{
		s: "不到叶子节点就找到了 targetSum, false", node: []interface{}{1, 2, nil, 3, nil, 4, nil, 5},
		target: 6, valid: false,
	},
}

func Test_HasPathSum(t *testing.T) {
	for _, mock := range MockDatas {
		t.Run(mock.s, func(t *testing.T) {
			node, err := generator(mock.node)
			NewWithT(t).Expect(err).To(BeNil())

			ok := hasPathSum(node, mock.target)
			NewWithT(t).Expect(ok).To(Equal(mock.valid))
		})
	}
}
