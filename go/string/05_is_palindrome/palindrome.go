package ispalindrome

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	// 字符判定
	// 0P 自己的代码为什么会是 true。
	// 因为这里只考虑了字母， 没有考虑数字。 所以 i 在遇到 0 的时候向后移动了
	//    re := regexp.MustCompile(`[a-z]`)
	re := regexp.MustCompile(`[a-z0-9]`)

	// 双指针

	// 外层大循环， 左右移动
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {

		// 小循环, 直到 s[i] 为字母
		for !re.Match([]byte{s[i]}) && i < j {
			i++
		}

		for !re.Match([]byte{s[j]}) && j > i {
			j--
		}

		// 如果两指针相遇
		if i == j {
			return true
		}

		// 如果两个都为字符
		// 1. 如果，不等
		if s[i] != s[j] {
			return false
		}

		// 2. 下一位，继续循环
	}

	return true
}

// // 结果很差
// 执行结果：
// 通过
// 显示详情
// 执行用时：
// 12 ms
// , 在所有 Go 提交中击败了
// 19.53%
// 的用户
// 内存消耗：
// 3.4 MB
// , 在所有 Go 提交中击败了
// 22.42%
// 的用户
