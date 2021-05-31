/*
移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。
相关标签
数组
双指针

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2ba4i/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

package movezeroes

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

// moveZeroes 20200531
func moveZeroes(nums []int) {

	// 使用双循环
	for zeroptr := 0; zeroptr < len(nums); zeroptr++ {
		// 如果 zeroptr 位置值不为 0, 继续下次循环
		if nums[zeroptr] != 0 {
			continue
		}

		// 从最后一个 0 位置开始， 找到第一个不为 0
		// for i:=zeroptr;i<len(nums)i++{}
		for i := zeroptr; i < len(nums); i++ {
			if nums[i] == 0 {
				continue
			}

			// 交换值
			nums[zeroptr], nums[i] = nums[i], nums[zeroptr]

			// 跳出， 并寻找最新的一个 0
			break

		}
	}
}

func Test_MoveZeroes(t *testing.T) {

	datas := []struct {
		result []int
		data   []int
	}{
		{
			result: []int{1, 3, 12, 0, 0},
			data:   []int{0, 1, 0, 3, 12},
		},
		{
			result: []int{1, 0, 0, 0, 0},
			data:   []int{0, 0, 0, 0, 1},
		},
		{
			result: []int{1, 0},
			data:   []int{1, 0},
		},
	}

	for _, data := range datas {
		t.Run(fmt.Sprint(data), func(t *testing.T) {
			moveZeroes(data.data)
			gomega.NewWithT(t).Expect(data.data).To(gomega.Equal(data.result))
		})
	}
}
