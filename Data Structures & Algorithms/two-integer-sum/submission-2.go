func twoSum(nums []int, target int) []int {
	result := []int{-1, -1}
	idx_map := make(map[int]int)
	for idx, v := range nums {
		idx_map[v] = idx
	}
	for idx, v := range nums {
		val, ok := idx_map[target-v]
		if ok && idx!= val {
			result[0] = idx
			result[1] = val
			return result
		}
	}
	return result
}

