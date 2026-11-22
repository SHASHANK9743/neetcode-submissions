func productExceptSelf(nums []int) []int {
	result := []int{}
	prefix_product := make([]int, len(nums))
	suffix_product := make([]int, len(nums))
	for idx, v := range nums {
		idx_end := len(nums) - idx - 1
		if idx == 0 {
			prefix_product[idx] = v
			suffix_product[idx_end] = nums[idx_end]
		} else {
			prefix_product[idx] = v * prefix_product[idx-1]
			suffix_product[idx_end] = nums[idx_end] * suffix_product[idx_end+1]
		}
	}
	for idx, _ := range nums {
		if idx == 0 {
			result = append(result, suffix_product[idx+1])
		} else if idx == len(nums)-1 {
			result = append(result, prefix_product[idx-1])
		} else {
			result = append(result, prefix_product[idx-1]*suffix_product[idx+1])
		}
	}
	return result
}