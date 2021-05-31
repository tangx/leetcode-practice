/*
两个数组的交集 II
给定两个数组，编写一个函数来计算它们的交集。



示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]
示例 2:

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]


说明：

输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
我们可以不考虑输出结果的顺序。
进阶：

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小很多，哪种方法更优？
如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
相关标签
排序
哈希表
双指针
二分查找

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2y0c2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
package intersect

import (
	"fmt"
	"sort"
	"testing"

	. "github.com/onsi/gomega"
)

func intersect(nums1 []int, nums2 []int) []int {

	return doublePointer(nums1, nums2)
}

// 双指针
func doublePointer(nums1, nums2 []int) []int {
	// 无论如何， 优先排序， 可以减少指针移动次数。 时间 O1
	sort.Ints(nums1)
	sort.Ints(nums2)

	result := []int{}

	// i 是 nums1 的指针位置
	// j 是 nums2 的指针位置
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		// 如果两个相等， 增加入结果，并都后移一位。
		// 要求: 输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
		// 因此每一次相等都要添加
		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i++
			j++
			continue
		}

		// 谁小移动谁
		if nums1[i] < nums2[j] {
			i++
			continue
		}
		// 谁小移动谁
		if nums1[i] > nums2[j] {
			j++
		}
	}
	return result
}

func Test_Intersect(t *testing.T) {
	datas := []struct {
		Result []int
		Nums1  []int
		Nums2  []int
	}{
		{
			Result: []int{2, 2},
			Nums1:  []int{1, 2, 2, 1},
			Nums2:  []int{2, 2},
		},
		{
			Result: []int{4, 9},
			Nums1:  []int{4, 9, 5},
			Nums2:  []int{9, 4, 9, 8, 4},
		},
	}

	for _, data := range datas {
		t.Run("DoublePoint"+fmt.Sprint(data), func(t *testing.T) {
			ret := doublePointer(data.Nums1, data.Nums2)
			NewWithT(t).Expect(ret).To(Equal(data.Result))
		})
	}

}
