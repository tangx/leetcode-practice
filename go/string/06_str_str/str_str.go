package strstr

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	// hello => 5
	//    ll => 2
	// 0-4
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		// 首字母相同
		if haystack[i] == needle[0] {

			// 原字符串上取相同长度的切片
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
		if haystack[i] == needle[0] && haystack[i:i+len(needle)] == needle {
			return i
		}

		continue
	}

	return -1
}
