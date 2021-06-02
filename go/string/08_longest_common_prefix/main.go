package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	// 边界判断

	// 实验组数量
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	// 如果有一个数组长度为0， 则 pre=""
	for _, str := range strs {
		if len(str) == 0 {
			return ""
		}
	}

	// ref 参照组, exp 实验组
	ref, exps := strs[:1], strs[1:]
	// 参照组字符串
	refstr := ref[0]

	for i := 1; i <= len(refstr); i++ {

		pre := refstr[:i]
		// fmt.Println("pre=", pre)

		for _, str := range exps {
			// fmt.Printf("str=%s, v=%s \n", str, str[:i])

			// 部分匹配
			// pre 长度超str // ["ab", "a"]
			if len(pre) > len(str) || pre != str[:i] {
				return pre[:i-1]
			}

		}
	}

	// 全部匹配
	return refstr
}
