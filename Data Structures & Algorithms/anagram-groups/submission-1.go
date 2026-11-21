func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	pos_map := make(map[string][]string)
	for _, itm := range strs {
		temp := []rune(itm)
		sort.Slice(temp, func(i, j int) bool {
			return temp[i] < temp[j]
		})
		temp_str := string(temp)
		_, ok := pos_map[temp_str]
		if ok {
			pos_map[temp_str] = append(pos_map[temp_str], itm)
		} else {
			pos_map[temp_str] = []string{itm}
		}
	}
	for _, v := range pos_map {
		result = append(result, v)
	}
	return result
}
