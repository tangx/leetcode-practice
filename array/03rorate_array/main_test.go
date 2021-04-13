package rotate03

import (
	"fmt"
	"testing"
)

/*

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

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2skh7/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


*/

func Test_rotate(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	// arr := []int{1}
	k := 3

	rotate(arr, k)
	fmt.Println(arr)

}

/*
解:
	虽然 slice 是引用类型， 使用 nums[i] 索引方式， 可以修改底层数据值。
	但是， append() 函数会生成一个新的 底层数组B，不会影响原有数组A。 因此函数内 nums 指针指向的底层数组变脸，但在函数外，底层数据A 的数据依旧不变。

	https://zhuanlan.zhihu.com/p/51203586
*/
func rotate(nums []int, k int) {

	n := len(nums)
	k = k % n
	num1 := nums[:n-k]
	num2 := nums[n-k:]

	// fmt.Println(num1, num2)
	numNew := append(num2, num1...)

	// 这里需要使用 copy 是因为 append 会生成一个新的底层数组。
	// 即使 nums=append() 这样也是只是 nums 变量指向新底层数组的地址。
	copy(nums, numNew)

}
