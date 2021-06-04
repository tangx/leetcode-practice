/*
对称二叉树
给定一个二叉树，检查它是否是镜像对称的。



例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3


但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3


进阶：

你可以运用递归和迭代两种方法解决这个问题吗？

相关标签
树
深度优先搜索
广度优先搜索

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/data-structure-binary-tree/xoxzgv/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
package issymmetric

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

// interface{} 中的元素， 数字与字母的判断
func Test_IntEqualNil(t *testing.T) {
	arr := []interface{}{3, nil, 3, "3"}

	for i := 1; i < len(arr); i++ {
		t.Run(fmt.Sprint(arr[0], arr[i]), func(t *testing.T) {
			gomega.NewWithT(t).Expect(arr[0]).To(gomega.Equal(arr[i]))
		})
	}
}
