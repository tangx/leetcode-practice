/*
加一
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。



示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
示例 3：

输入：digits = [0]
输出：[1]


提示：

1 <= digits.length <= 100
0 <= digits[i] <= 9
相关标签


作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2cv1c/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

package plusone

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

func plusOne(digits []int) []int {
	n := len(digits)
	for ; n > 0; n-- {
		// 不等于 9 ，不用进位， 直接 +1 并返回
		if digits[n-1] != 9 {
			digits[n-1] += 1
			return digits
		}

		// 如果等于9, 本位设置为 0 ， 进位进行下一次循环
		digits[n-1] = 0
	}

	// 当执行到此处时， 表示所有位都位 9， 需要返回进位的新数组。
	digits = append([]int{1}, digits...)

	return digits
}

func Test_PlusOne(t *testing.T) {
	datas := []struct {
		Result []int
		Data   []int
	}{
		{
			Result: []int{1, 2, 4},
			Data:   []int{1, 2, 3},
		},
		{
			Result: []int{1, 0, 0},
			Data:   []int{9, 9},
		},
		{
			Result: []int{9, 0},
			Data:   []int{8, 9},
		},
	}

	for _, data := range datas {
		t.Run(fmt.Sprint(data), func(t *testing.T) {
			ret := plusOne(data.Data)
			gomega.NewWithT(t).Expect(ret).To(gomega.Equal(data.Result))
		})
	}
}
