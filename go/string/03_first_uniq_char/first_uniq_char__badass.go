package firstuniqchar

func firstUniqChar_badass(s string) int {

	// hset 只保存出现过一次的
	hset := make(map[rune]int)
	// hsetdup 保存重复出现的字符
	dupset := make(map[rune]struct{})

	for idx, value := range s {
		// 三进宫， 跳过
		if _, ok := dupset[value]; ok {
			continue
		}
		// 二进宫， 转移
		if _, ok := hset[value]; ok {
			// 出现2次以上， 添加删除添加删除
			delete(hset, value)
			dupset[value] = struct{}{}
			continue
		}

		// 一进宫， 上位
		hset[value] = idx
	}

	// 如果不存在孤独字符，则返回 -1
	if len(hset) == 0 {
		return -1
	}

	// 使用最大值
	// 由于 0 不比任何其他值大， 因此这里不能使用 0 作为 pos 的默认值。
	pos := len(s)
	for _, idx := range hset {
		if idx < pos {
			pos = idx
		}
	}

	return pos
}
