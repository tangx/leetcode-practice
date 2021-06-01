/*
字符串中的第一个唯一字符
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。



示例：

s = "leetcode"
返回 0

s = "loveleetcode"
返回 2


提示：你可以假定该字符串只包含小写字母。

相关标签
哈希表
字符串

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn5z8r/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
package firstuniqchar

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

var MockDatas = []struct {
	Data   string
	Result int
}{
	{Data: "leetcode", Result: 0},
	{Data: "loveleetcode", Result: 2},
	{Data: "s", Result: 0},
	{Data: "ss", Result: -1},
	{Data: "aadadaad", Result: -1},
}

func Test_FirstUniqChar(t *testing.T) {

	f := func(fn func(string) int) {
		for _, mock := range MockDatas {
			t.Run(fmt.Sprint(mock), func(t *testing.T) {
				ret := fn(mock.Data)
				gomega.NewWithT(t).Expect(ret).To(gomega.Equal(mock.Result))
			})
		}
	}

	f(firstUniqChar_badass)
	f(firstUniqChar)
}
