package reverseint32

import (
	"math"
	"strconv"
)

// reverse_int32_20210601 翻转一个有符号数字
// 1. 转成字符串
// 2. 转成 bytes
// 3. 使用双指针获取新字符串 b2
// 4. 分离符号与数字， 翻转数字
// 5. 组合符号与数字， 转换为 int
func reverse_int32_20210601(x int) int {

	// 处理边界问题
	if x := borderHandler(x); x == 0 {
		return 0
	}

	// 1. 2.
	s := strconv.Itoa(x)
	b := []byte(s)

	// 3. 双指针， 去尾数的 0
	//   1. fast 向后， 如遇到0 继续移动， 如果非0
	//   2. slow 移动到 fast 位置
	//   3. 取 0 到 slow 的数字。 注意边界
	tail := 0
	for i := 0; i < len(b); i++ {
		if b[i] == '0' {
			continue
		}

		tail = i
	}

	// 4. 新字节翻转
	// 4.1 新字节
	negative := false
	b2 := b[:tail+1]
	if b[0] == '-' {
		negative = true
		b2 = b2[1:]
	}

	// 4.2 前后翻转
	// 这里使用 b2 不包含符号， 在翻转的时候就方便多了
	nb2 := len(b2)
	for i := 0; i < nb2/2; i++ {
		b2[i], b2[nb2-1-i] = b2[nb2-1-i], b2[i]
	}

	if negative {
		b2 = append([]byte{'-'}, b2...)
	}
	s = string(b2)
	x, _ = strconv.Atoi(s)

	// 2^31 在 golang 中是 **位运算** 是异或。
	// 2^31 = 29
	return borderHandler(x)
}

func borderHandler(x int) int {
	border := int(math.Pow(2, 31))

	if x < 0-border || x > border-1 {
		return 0
	}
	return x
}
