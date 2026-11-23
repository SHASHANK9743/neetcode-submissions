func twoSum(numbers []int, target int) []int {
	result := []int{-1, -1}
	for idx, v := range numbers {
		diff := target - v
		if diff > numbers[len(numbers)-1] {
			continue
		}
		low := idx + 1
		high := len(numbers) - 1
		for low <= high {
			mid := (low + high) / 2
			if diff == numbers[mid] {
				result[0] = idx + 1
				result[1] = mid + 1
				break
			} else if diff > numbers[mid] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return result
}
