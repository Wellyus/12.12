package str

func Str(a, b string) int {
	s1, s2 := []byte(a), []byte(b)
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] < s2[i] {
			return -1
		}
		if s1[i] > s2[i] {
			return 1
		}
	}
	if len(a) < len(b) {
		return -1
	}
	if len(a) > len(b) {
		return 1
	}
	return 0
}

//
