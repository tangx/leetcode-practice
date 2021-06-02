/*
最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。



示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。


提示：

0 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 仅由小写英文字母组成
相关标签
字符串

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnmav1/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
package longestcommonprefix

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

var MockDatas = []struct {
	strs []string
	pre  string
}{
	{strs: []string{""}, pre: ""},
	{strs: []string{"abc"}, pre: "abc"},
	{strs: []string{"flower", "flow", "flight"}, pre: "fl"},
	{strs: []string{"dog", "racecar", "car", "dot"}, pre: ""},
	{strs: []string{"flower", "flower"}, pre: "flower"},
	// 长度导致 out of range
	{strs: []string{"ab", "a"}, pre: "a"},
}

func Test_LongestCommonPrefix(t *testing.T) {
	for _, mock := range MockDatas {
		t.Run(fmt.Sprint(mock), func(t *testing.T) {
			pre := longestCommonPrefix(mock.strs)
			gomega.NewWithT(t).Expect(pre).To(gomega.Equal(mock.pre))
		})
	}
}
