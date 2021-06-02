/*整数反转
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。


示例 1：

输入：x = 123
输出：321
示例 2：

输入：x = -123
输出：-321
示例 3：

输入：x = 120
输出：21
示例 4：

输入：x = 0
输出：0


提示：

-231 <= x <= 231 - 1
相关标签
数学

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnx13t/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
package reverseint32

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

// var MockDatas = []struct {
// 	Data   int
// 	Result int
// }{
// 	{Data: 1, Result: 1},
// 	{Data: 120, Result: 21},
// 	{Data: -120, Result: -21},
// 	{Data: 1245677899, Result: 0},
// }

func Test_Reverse(t *testing.T) {

	f := func(fn func(int) int) {
		for _, mock := range MockDatas {
			t.Run(fmt.Sprint(mock), func(t *testing.T) {
				ret := fn(mock.Data)
				gomega.NewWithT(t).Expect(ret).To(gomega.Equal(mock.Result))
			})
		}
	}

	f(reverse)
}

func Test_NegativeMod(t *testing.T) {

	f := func(x int) {
		fmt.Printf("%d/10=%d\n", x, x/10)
		fmt.Printf("%d mod 10=%d\n", x, x%10)
	}

	f(11)
	f(-11)
}
