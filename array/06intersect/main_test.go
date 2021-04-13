package intersect06

import (
	"fmt"
	"sort"
	"testing"
)

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


作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2y0c2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func Test_intersect(t *testing.T) {
	// nums1 := []int{1, 2, 2, 1}
	// nums2 := []int{2, 2}

	// nums1 := []int{4, 9, 5}
	// nums2 := []int{9, 4, 9, 8, 4}

	nums1 := []int{1, 2}
	nums2 := []int{1, 1}

	ret := intersect(nums1, nums2)
	fmt.Println(ret)
}

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	result := []int{}

	var i int = 0
	var j int = 0

	for i < len(nums1) && j < len(nums2) {

		switch {
		case nums1[i] == nums2[j]:
			result = append(result, nums1[i])
			// nums1, nums2 指针都向右移动一位
			j++
			i++

		case nums1[i] > nums2[j]:
			// nums2 指针向右移动一位
			j++
		case nums1[i] < nums2[j]:
			// nums1 指针向右移动一位
			// continue
			i++
		}
	}

	return result
}

func Test_LoopGoto(t *testing.T) {

	var i int = 0
	var j int = 0
	for ; i < 10; i++ {
		for ; j < 10; j++ {
			if i == j {
				fmt.Printf("breaking i=>%d, j=>%d \n", i, j)
				goto outbreak

			}

			fmt.Printf("=======> i=>%d, j=>%d \n", i, j)
		}
	outbreak:
		j += 1
	}
}
