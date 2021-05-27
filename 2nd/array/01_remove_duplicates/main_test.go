package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

/*

删除排序数组中的重复项
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。


说明:

为什么返回数值是整数，但输出的答案是数组呢?

请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

你可以想象内部操作如下:

// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
int len = removeDuplicates(nums);

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}

示例 1：

输入：nums = [1,1,2]
输出：2, nums = [1,2]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
示例 2：

输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。


提示：

0 <= nums.length <= 3 * 104
-104 <= nums[i] <= 104
nums 已按升序排列


作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2gy9m/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

// removeDuplicates 删除重复值
// 解题思路， 使用双指针
func removeDuplicates(nums []int) int {

	// 如果长度为 0
	// nil 长度也为0
	if len(nums) == 0 {
		return 0
	}

	// 如果长度不为0
	head := 0

	// 移动指针 tail
	for tail := 1; tail < len(nums); tail++ {
		// 如果 nums[tail] 与 nums[head] 相同， 则 tail 继续移动，
		if nums[tail] == nums[head] {
			continue
		}

		// 如果 nums[tail] 与 nums[head] 不相同， 则 head+1 并复制 nums[tail]
		head += 1
		nums[head] = nums[tail]
	}

	// 返回长度 head+1
	// 因为 head 指针位置从 0 开始， 因此长度要 +1
	return head + 1
}

func Test_removeDuplicates(t *testing.T) {

	// 常规操作，先去除异常场景。把数组是null和数组为空的情况排除。
	t.Run("nil slice", func(t *testing.T) {
		n := removeDuplicates(nil)
		NewWithT(t).Expect(n).To(Equal(0))
	})
	t.Run("len = 0", func(t *testing.T) {
		nums := []int{}
		n := removeDuplicates(nums)
		NewWithT(t).Expect(n).To(Equal(0))
	})

	// 正常判断
	t.Run("len = 1", func(t *testing.T) {
		nums := []int{1}
		n := removeDuplicates(nums)
		NewWithT(t).Expect(n).To(Equal(1))
	})

	t.Run("len > 2: [1,1,2]", func(t *testing.T) {
		nums := []int{1, 1, 2}
		n := removeDuplicates(nums)
		NewWithT(t).Expect(n).To(Equal(2))
	})

	t.Run("len > 2: [0,0,1,1,1,2,2,3,3,4]", func(t *testing.T) {
		nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		n := removeDuplicates(nums)
		NewWithT(t).Expect(n).To(Equal(5))
	})
}
