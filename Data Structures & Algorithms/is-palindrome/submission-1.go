func isPalindrome(s string) bool {
	s_bytes := []int{}
	for _, v := range s {
		if (v > 64 && v < 91) || (v > 96 && v < 123) || (v > 47 && v < 58) {
			if v > 96 && v < 123 {
				v = v - 32
			}
			s_bytes = append(s_bytes, int(v))
		}
	}
	i := 0
	j := len(s_bytes) - 1
	for i <= j {
		if i == j {
			return true
		} else {
			if s_bytes[i] != s_bytes[j] {
				return false
			}
		}
		j = j - 1
		i = i + 1
	}
	return true
}
