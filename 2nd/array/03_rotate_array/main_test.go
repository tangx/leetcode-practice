/*

旋转数组
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。



进阶：

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？


示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]


提示：

1 <= nums.length <= 2 * 104
-231 <= nums[i] <= 231 - 1
0 <= k <= 105
相关标签
数组

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2skh7/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

// rotate 执行 原地算法
// 原地算法: 空间 O1, 不增加额外内存开销保存拆分的数据
// 使用循环不断移动数字到指定位置
// 因为 nums 是 slice， 所以在函数内修改也会影响外部
func rotate(nums []int, k int) {

	if len(nums) < 2 {
		return
	}

	// /* 移动方法 1 */
	// step := k % len(nums) : k > len(nums) 不用进行多次循环
	// for step := k % len(nums); step > 0; step-- {
	// 	// 将最后一个值取出来暂存
	// 	tmp := nums[len(nums)-1]
	// 	// 每一位向后移动
	// 	for i := len(nums) - 1; i >= 0; i-- {

	// 		// 如果已经后移到最后第一个位置
	// 		// 将暂存值放入第一个位置
	// 		if i == 0 {
	// 			nums[0] = tmp
	// 			continue
	// 		}
	// 		// 后移
	// 		nums[i] = nums[i-1]
	// 	}
	// }

	/* 移动方法2
	测试用例超时， 继续优化
	https://leetcode-cn.com/submissions/detail/181609182/
	*/
	// 取余， 如果 k > len(nums) 也不用进行额外循环
	// for step := k % len(nums); step > 0; step-- {
	// 	for i := 0; i < len(nums); i++ {
	// 		// 位置交换
	// 		// 每次交换一个位置， 将 nums[0] 作为暂存区
	// 		nums[0], nums[i] = nums[i], nums[0]
	// 	}
	// }

}

func Test_Rotate(t *testing.T) {

	datas := []struct {
		Key    int
		Origin []int
		Result []int
	}{
		{0, []int{}, []int{}},
		{1, []int{1}, []int{1}},
		{3, []int{1, 2, 3, 4, 5, 6, 7}, []int{5, 6, 7, 1, 2, 3, 4}},
		{2, []int{-1, -100, 3, 99}, []int{3, 99, -1, -100}},
	}

	t.Run("normal", func(t *testing.T) {

		for _, data := range datas {
			rotate(data.Origin, data.Key)
			NewWithT(t).Expect(data.Result).To(Equal(data.Origin))
		}

	})
}
