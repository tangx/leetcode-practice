/*
存在重复元素
给定一个整数数组，判断是否存在重复元素。

如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。



示例 1:

输入: [1,2,3,1]
输出: true
示例 2:

输入: [1,2,3,4]
输出: false
示例 3:

输入: [1,1,1,3,3,4,3,2,4,2]
输出: true
相关标签
数组
哈希表

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x248f5/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
package main

import (
	"sort"
	"testing"

	. "github.com/onsi/gomega"
)

// containsDuplicateByHashset 判断数组是否包含重复元素
// 使用 hash 桶 （map） 进行封装
// 优点: 不会篡改原数组
// 缺点: 需要开辟新内存保存 hash桶
func containsDuplicateByHashset(nums []int) bool {
	// map的value值是布尔型，这会导致set多占用内存空间，解决这个问题，则可以将其替换为空结构。在Go中，空结构通常不使用任何内存。
	// tank := make(map[int]bool)
	tank := make(map[int]struct{})

	for _, i := range nums {
		if _, ok := tank[i]; ok {
			return true
		}
		tank[i] = struct{}{}
	}
	return false
}

// containsDuplicateBySort 先排序， 在相邻比较
// 优点: 不开辟新内存
// 缺点: 篡改原数组
func containsDuplicateBySort(nums []int) bool {
	sort.Ints(nums)

	// 这里需要注意数组边界越界
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func containsDuplicate(nums []int) bool {
	return containsDuplicateBySort(nums)
}

func Test_ContainsDuplicate(t *testing.T) {

	datas := []struct {
		IsDup bool
		Data  []int
	}{
		{IsDup: true, Data: []int{1, 2, 3, 1}},
		{IsDup: false, Data: []int{1, 2, 3, 4}},
		{IsDup: true, Data: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
	}

	t.Run("ContainsDuplicate", func(t *testing.T) {
		for _, data := range datas {
			NewWithT(t).Expect(containsDuplicate(data.Data)).To(Equal(data.IsDup))
		}
	})
}
