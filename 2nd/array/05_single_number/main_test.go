/*
只出现一次的数字
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
示例 2:

输入: [4,1,2,1,2]
输出: 4
相关标签
位运算
哈希表

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x21ib6/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

package singlenumber

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func singleNumber(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	return bitOperation(nums)
}

// 位运算
func bitOperation(nums []int) int {
	// 1. 相同的数字进行 异或运算 结果位0 : a^a=0
	// 2. 异或运算 支持交换律 a^b^c=a^c^b=b^a^c
	// 3. 0 与任何数字 x 进行 异或运算，结果为 x: 0^x=x

	ret := 0

	for _, num := range nums {
		ret ^= num
	}

	return ret
}

func Test_BitOperation(t *testing.T) {
	fmt.Println(0 ^ 2)
	fmt.Println(2 ^ 2)
	fmt.Println(2 ^ 2 ^ 3)
}

// hash 桶
func hashset(nums []int) int {

	// 创建一个 hashset ， 保存元素
	set := make(map[int]struct{})

	for _, num := range nums {

		// 如果存在，配对成功。删除
		if _, ok := set[num]; ok {
			delete(set, num)
			continue
		}
		// 如果不存在， 则添加到 set 中
		set[num] = struct{}{}

	}

	// 因为只有一个，直接返回。
	for num := range set {
		return num
	}

	return 0
}

func Test_SingleNumber(t *testing.T) {

	datas := []struct {
		Result int
		Data   []int
	}{
		{Result: 1, Data: []int{2, 2, 1}},
		{Result: 4, Data: []int{4, 1, 2, 1, 2}},
	}

	for _, data := range datas {
		t.Run("HashSet"+fmt.Sprint(data), func(t *testing.T) {
			ret := hashset(data.Data)
			NewWithT(t).Expect(ret).To(Equal(data.Result))
		})
		t.Run("Bit Operation"+fmt.Sprint(data), func(t *testing.T) {
			ret := bitOperation(data.Data)
			NewWithT(t).Expect(ret).To(Equal(data.Result))
		})
	}

}
