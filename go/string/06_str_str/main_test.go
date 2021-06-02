/*
实现 strStr()
实现 strStr() 函数。

给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。



说明：

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。



示例 1：

输入：haystack = "hello", needle = "ll"
输出：2
示例 2：

输入：haystack = "aaaaa", needle = "bba"
输出：-1
示例 3：

输入：haystack = "", needle = ""
输出：0


提示：

0 <= haystack.length, needle.length <= 5 * 104
haystack 和 needle 仅由小写英文字符组成
相关标签
双指针
字符串

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnr003/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
package strstr

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

var MockDatas = []struct {
	haystack string
	needle   string
	index    int
}{
	// {haystack: "hello", needle: "ll", index: 2},
	// {haystack: "aaaaa", needle: "bba", index: -1},
	// {haystack: "", needle: "", index: 0},
	{haystack: "a", needle: "a", index: 0},
}

func Test_StrStr(t *testing.T) {

	for _, mock := range MockDatas {
		t.Run(fmt.Sprint(mock), func(t *testing.T) {
			i := strStr(mock.haystack, mock.needle)
			gomega.NewWithT(t).Expect(i).To(gomega.Equal(mock.index))
		})
	}
}
