package isanagram

import (
	"sort"
	"strings"
)

// isAnagram__sort 排序并按位对比
// 这道使用排序并不好， s/t 长度越长，排序越耗时。
func isAnagram__sort(s string, t string) bool {
	// 先对比长度
	if len(s) != len(t) {
		return false
	}

	sl := splitSort(s)
	tl := splitSort(t)

	// 对比
	// 字母异位词
	// 1. 所有出现字母次数必须一致
	// 2. 仅是排列顺序不一致
	for i := 0; i < len(sl); i++ {
		if sl[i] != tl[i] {
			return false
		}
	}
	return true
}

// splitSort 分裂并排序
func splitSort(s string) (l []string) {
	l = strings.Split(s, "")
	sort.Strings(l)
	return
}
