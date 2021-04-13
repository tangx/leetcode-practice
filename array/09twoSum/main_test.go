package twosum09

import (
	"fmt"
	"testing"
)

/*
两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。



示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]


提示：

2 <= nums.length <= 103
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2jrse/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func Test_TwoSum(t *testing.T) {

	for target, nums := range map[int][]int{
		3: {4, 3, 1, 2},
		9: {2, 7, 11, 15},
	} {
		r := twoSum(nums, target)
		fmt.Println(r)
	}
}

// 解题思路， 创建一个 map[val]pos // map[差]索引
// 如果 target - nums[i] 不在 map 中， 则继续。
// 否则返回位置
func twoSum(nums []int, target int) []int {

	m := make(map[int]int)
	for idx, val := range nums {
		if pos, ok := m[val]; ok {
			return []int{pos, idx}
		}

		// 保存所需的 差 为key。 差即为后续 数列中的值。
		m[target-val] = idx

	}

	return nil
}
