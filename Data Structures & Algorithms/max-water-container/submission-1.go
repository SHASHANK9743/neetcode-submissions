func maxArea(heights []int) int {
	i := 0
	j := len(heights) - 1
	max_area := -1
	for (j - i) > 0 {
		area := min(heights[i], heights[j]) * (j - i)
		if area > max_area {
			max_area = area
		}
		if heights[i] <= heights[j] {
			i = i + 1
		} else {
			j = j - 1
		}
	}
	return max_area
}