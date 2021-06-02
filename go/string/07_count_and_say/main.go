package countandsay

import (
	"bytes"
	"strconv"
)

func countAndSay(n int) string {
	if n < 1 || n > 30 {
		return ""
	}

	s := "1"

	// i 不能从 0 开始， 因为边界是 0<n<31
	for i := 1; i < n; i++ {
		s = countString(s)
	}

	return s
}

func countString(s string) string {
	// 双指针
	buf := bytes.NewBuffer(nil)

	// 11122211
	// i  j        // 0,3
	//    i  j     // 3,6
	//       i j   // 6,8
	for i, j := 0, 0; i < len(s); {

		// j 向右移动,
		// 直到 超出边界或字面值不等
		for ; j < len(s) && s[i] == s[j]; j++ {
			continue
		}

		// 字面值不等

		// n 字面值 x 出现的次数
		n := strconv.Itoa(j - i)
		// fmt.Printf("j-i=n => %d - %d = %s \n", j, i, n)
		// 字面量 nx
		buf.WriteString(n)
		buf.WriteByte(s[i])

		// 移动 i 到最新位置
		i = j
	}

	return buf.String()
}
