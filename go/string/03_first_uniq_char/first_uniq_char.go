package firstuniqchar

// firstUniqChar 双循环
// 一循环， 统计出现次数并保存到 hset 中
// 二循环， 从 hset 中返回第一次出现的字符
func firstUniqChar(s string) int {

	hset := make(map[rune]int)

	for _, char := range s {
		hset[char] += 1
	}

	for idx, char := range s {
		if hset[char] == 1 {
			return idx
		}
	}

	return -1
}
