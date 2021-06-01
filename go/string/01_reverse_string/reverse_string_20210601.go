package reversestring

func reverseString_20210601(s []byte) {
	edge := len(s) - 1
	for i := 0; i <= edge/2; i++ {
		s[i], s[edge-i] = s[edge-i], s[i]
	}
}
