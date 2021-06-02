package reverseint32

// reverse 翻转数字
func reverse(x int) int {
	// 如果 x 不满足 int32， 则返回0
	if !validateInt32(x) {
		return 0
	}

	rex := 0
	for ; x != 0; x /= 10 {
		rex = rex*10 + x%10
		if !validateInt32(rex) {
			return 0
		}
	}

	return rex
}

// validateInt32 判断数字是否在 int32 范围内
func validateInt32(x int) bool {
	i32 := int32(x)
	return int(i32) == x
}
