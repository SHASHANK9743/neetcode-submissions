func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	max_len := -1
	unique_map := make(map[int]bool)
	for _, v := range nums {
		unique_map[v] = true
	}
	starting_map := make(map[int]bool)
	for _, v := range nums {
		_, ok := unique_map[v-1]
		if !ok {
			starting_map[v] = true
		}
	}
	for k, _ := range starting_map {
		curr_val := k
		exists := true
		curr_len := 1
		for exists {
			_, ok := unique_map[curr_val+1]
			if ok {
				curr_len = curr_len + 1
				curr_val = curr_val + 1
			} else {
				exists = false
				if curr_len > max_len {
					max_len = curr_len
				}
			}
		}
	}
	return max_len
}