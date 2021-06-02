package isanagram

func isAnagram__loopOnce(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	set := make(map[byte]int)

	counter := 0
	for i := 0; i < len(s); i++ {
		sv := s[i]
		if set[sv] += 1; set[sv] == 1 {
			counter++
		}
		tv := t[i]
		if set[tv] -= 1; set[tv] == 0 {
			counter--
		}
	}

	return counter == 0
}
