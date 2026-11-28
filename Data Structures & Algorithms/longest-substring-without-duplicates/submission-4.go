func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	max_len := 1
	i := 0
	j := i + 1
	freq_map := make(map[string]int)
	freq_map[string(s[i])] = 1
	for i < len(s) && j < len(s) {
		_, ok := freq_map[string(s[j])]
		if ok {
			if string(s[i]) == string(s[j]) {
				j = j + 1
				i = i + 1
			} else {
				delete(freq_map, string(s[i]))
				i = i + 1
			}
		} else {
			fmt.Println(string(s[i : j+1]))
			current_length := j - i + 1
			if current_length > max_len {
				max_len = current_length
			}
			freq_map[string(s[j])] = freq_map[string(s[j])] + 1
			j = j + 1
		}
	}
	return max_len
}