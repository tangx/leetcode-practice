package myatoi

import (
	"regexp"
	"strings"
)

func myAtoi(s string) int {
	// 去除左右空格
	s = strings.TrimSpace(s)

	// 使用正则查找匹配内容
	re := regexp.MustCompile(`^([+-]?)([0-9]*)`)

	// parts=[all,symbol,value]
	// parts=[-123,-,123]
	// parts=[123,,123]
	// parts=[,,]
	parts := re.FindStringSubmatch(s)

	sum := 0

	MAX_VALUE := 1 << 31
	for i := 0; i < len(parts[2]) && sum < MAX_VALUE; i++ {
		sum = sum*10 + int(parts[2][i]-'0')
	}
	if parts[1] == "-" {
		sum = 0 - sum
	}
	return cutoff(sum)
}

// cutoff 使用位运算进行边界判断
// https://play.golang.org/p/VEH1er0D_3p
// 直接写， 不进行位运算， 可能性能要好点
func cutoff(n int) int {
	switch {
	case n > 2147483647:
		// 1<<31-1
		return 2147483647
	case n < -2147483648:
		// -1<<31
		return -2147483648
	default:
		return n
	}
}

// 这里是截断， 因此不能直接转
// func i32(n int) int32 {
// 	return int32(n)
// }
