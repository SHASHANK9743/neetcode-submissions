func twoSum(inp_slice []int, target int) [][]int {
	result := [][]int{}
	unique_map := make(map[int]int)

	for idx, v := range inp_slice {
		unique_map[v] = idx
	}

	for idx, v := range inp_slice {
		val, ok := unique_map[target-v]
		if ok && val != idx {
			temp_result := []int{}
			temp_result = append(temp_result, v)
			temp_result = append(temp_result, target-v)
			result = append(result, temp_result)
		}
	}
	return result
}

func threeSum(nums []int) [][]int {
	result := [][]int{}
	seen := make(map[[3]int]bool)
	sort.Ints(nums)
	for idx, v := range nums {
		temp_slice := []int{}
		temp_slice = append(temp_slice, nums[idx+1:]...)
		target := -1 * v
		two_sum_results := twoSum(temp_slice, target)
		if len(two_sum_results) > 0 {
			for _, two_sum := range two_sum_results {
				sorted_result := []int{v, two_sum[0], two_sum[1]}
				sort.Ints(sorted_result)
				key := [3]int{sorted_result[0], sorted_result[1], sorted_result[2]}
				if !seen[key] {
					seen[key] = true
					result = append(result, sorted_result)
				}
			}
		}
	}
	return result
}