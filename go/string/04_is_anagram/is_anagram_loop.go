package isanagram

func isAnagram__loop(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	// 统计字母出现次数
	sset := make(map[byte]int)
	tset := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		v := s[i]
		sset[v] += 1

		v = t[i]
		tset[v] += 1

	}

	// 对比字母出现次数
	for v := range sset {
		if sset[v] != tset[v] {
			return false
		}
	}

	return true
}
