func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++ // a missing key starts at 0, so no ok-check needed
	}

	// buckets[c] holds the numbers that appear exactly c times
	buckets := make([][]int, len(nums)+1)
	for n, c := range freq {
		buckets[c] = append(buckets[c], n)
	}

	// walk from the highest count down until we have k numbers
	result := make([]int, 0, k)
	for c := len(buckets) - 1; c > 0 && len(result) < k; c-- {
		result = append(result, buckets[c]...)
	}
	return result[:k]
}

