func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1_freq_map := make(map[string]int)
	window_size := len(s1)
	for _, v := range s1 {
		_, ok := s1_freq_map[string(v)]
		if ok {
			s1_freq_map[string(v)] = s1_freq_map[string(v)] + 1
		} else {
			s1_freq_map[string(v)] = 1
		}
	}
	for i := 0; i < len(s2)-window_size+1; i = i + 1 {
		temp_string := string(s2[i : i+window_size])
		temp_freq_map := make(map[string]int)
		permutation_found := true
		for _, v := range temp_string {
			_, ok := temp_freq_map[string(v)]
			if ok {
				temp_freq_map[string(v)] = temp_freq_map[string(v)] + 1
			} else {
				temp_freq_map[string(v)] = 1
			}
		}
		for k, v := range s1_freq_map {
			freq, ok := temp_freq_map[k]
			if ok {
				if freq != v {
					permutation_found = false
					break
				}
			} else {
				permutation_found = false
				break
			}
		}
		if permutation_found {
			return true
		}
	}
	return false
}