package movezeros08

import (
	"fmt"
	"testing"
)

/*
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

func Test_MoreZero(t *testing.T) {

	for _, nums := range [][]int{
		{0, 1, 0, 3, 12},
		{2, 0, 0, 9, 15, 3, 0},
	} {
		moveZeroes(nums)
		fmt.Println(nums)
	}
}

func moveZeroes(nums []int) {
	// 需要移动 0 到的数组末尾位置
	head := 0
	tail := 0

	// head, tail 两个指针同时从左向右移动
	// 如果 head 为0， 则表示该位置需替换。
	// 如果 tail 不为0， 则表示该位置需要为 head 换位。
	for tail < len(nums) {
		switch {
		case nums[head] == 0 && nums[tail] != 0:
			nums[head], nums[tail] = nums[tail], nums[head]
			head++
			tail++
		case nums[head] != 0:
			head++
			fallthrough
		case nums[tail] == 0:
			tail++
		}
	}
}
