package plusone07

import (
	"fmt"
	"testing"
)

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
数组

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2cv1c/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func Test_plusOne(t *testing.T) {
	for _, digits := range [][]int{
		{1, 2, 3, 4},
		{1, 2, 9},
		{9, 9, 9},
	} {
		r := plusOne(digits)
		fmt.Println(r)
	}
}

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		// 当不是9时，不需要进位
		if digits[i] != 9 {
			digits[i] = digits[i] + 1
			return digits
		} else {
			// 进位, 本位为0， 进1 // 进入下一轮循环
			digits[i] = 0
		}
	}

	// 当所有值都为 9 时。 9999， 需要扩容
	ret := append([]int{1}, digits...)

	return ret
}
