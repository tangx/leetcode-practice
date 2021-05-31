/*
两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

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

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案
进阶：你可以想出一个时间复杂度小于 O(n2) 的算法吗？

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2jrse/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
package twosum

import (
	"fmt"
	"sort"
	"testing"

	"github.com/onsi/gomega"
)

func twoSum(nums []int, target int) []int {
	// 因为 元素不能重复出现在答案中， 即不能出现 4+4 = 8
	// 因此，即使在数据中出现 [4,4,2,6]的选项，可以将第二个 4 忽略。
	// set:=map[value]index
	set := make(map[int]int)

	for nindex, nvalue := range nums {

		/*
			但是，数组中同一个元素在答案里不能重复出现。

			nums=[4,4,2,6] , target=8

			答案中不能出现 [1,1] 这样的规则
			但可以出现 [1,2]

			理解错误
		*/

		// 如果存在相同 value ， 跳过
		// if _, ok := set[nvalue]; ok {
		// 	continue
		// }

		// 如果不存在， 则进行判断
		// 获取目标值
		value := target - nvalue
		// 判断是否存在， 存在并返回
		if index, ok := set[value]; ok {
			result := []int{index, nindex}

			// 排序， 为了方便写测试用例
			sort.Ints(result)
			return result
		}

		// 不存在，则添加
		set[nvalue] = nindex
	}

	return nil
}

func Test_TwoSum(t *testing.T) {
	datas := []struct {
		result []int
		data   []int
		target int
	}{
		{
			result: []int{0, 1},
			data:   []int{2, 7, 11, 15},
			target: 9,
		},
		{
			result: []int{1, 2},
			data:   []int{3, 2, 4},
			target: 6,
		},
		{
			result: []int{0, 1},
			data:   []int{4, 4, 2, 6, 7},
			target: 8,
		},
	}

	for _, data := range datas {
		t.Run(fmt.Sprint(data), func(t *testing.T) {
			ret := twoSum(data.data, data.target)
			gomega.NewWithT(t).Expect(ret).To(gomega.Equal(data.result))
		})
	}
}
