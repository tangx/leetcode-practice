package singlenumber

import (
	"fmt"
	"testing"
)

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


作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x21ib6/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func Test_singleNumber(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1}
	n := singleNumber(arr)
	fmt.Println(n)

	singleNumber2(arr)
}

func singleNumber(nums []int) int {

	m := make(map[int]bool)
	for _, val := range nums {
		if m[val] {
			delete(m, val)
			continue
		}
		m[val] = true
	}

	fmt.Println(m)
	for key := range m {
		return key
	}

	return 0
}

/*
答案是使用位运算。对于这道题，可使用异或运算 \oplus⊕。异或运算有以下三个性质。

任何数和 00 做异或运算，结果仍然是原来的数，即 a \oplus 0=aa⊕0=a。
任何数和其自身做异或运算，结果是 00，即 a \oplus a=0a⊕a=0。
异或运算满足交换律和结合律，即 a \oplus b \oplus a=b \oplus a \oplus a=b \oplus (a \oplus a)=b \oplus0=ba⊕b⊕a=b⊕a⊕a=b⊕(a⊕a)=b⊕0=b。


作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/single-number/solution/zhi-chu-xian-yi-ci-de-shu-zi-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/
func singleNumber2(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num

		fmt.Printf("%d ->(single) %d \n", num, single)
	}
	return single
}

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/single-number/solution/zhi-chu-xian-yi-ci-de-shu-zi-by-leetcode-solution/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
