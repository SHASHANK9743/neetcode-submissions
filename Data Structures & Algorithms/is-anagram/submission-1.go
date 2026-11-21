func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	} else {
		s_freq_map := make(map[string]int)
		t_freq_map := make(map[string]int)
		for _, char := range s {
			_, ok := s_freq_map[string(char)]
			if ok {
				s_freq_map[string(char)] = s_freq_map[string(char)] + 1
			} else {
				s_freq_map[string(char)] = 1
			}
		}
		for _, char := range t {
			_, ok := t_freq_map[string(char)]
			if ok {
				t_freq_map[string(char)] = t_freq_map[string(char)] + 1
			} else {
				t_freq_map[string(char)] = 1
			}
		}

		for k, v := range s_freq_map {
			count, ok := t_freq_map[k]
			if ok {
				if v != count {
					return false
				}
			} else {
				return false
			}
		}
		return true
	}
}
