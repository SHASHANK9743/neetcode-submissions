func maxProfit(prices []int) int {
	max_profit := 0
	max_so_far := make([]int, len(prices))
	j := len(prices) - 1
	for j > -1 {
		if j == len(prices)-1 {
			max_so_far[j] = prices[j]
		} else {
			if prices[j] > max_so_far[j+1] {
				max_so_far[j] = prices[j]
			} else {
				max_so_far[j] = max_so_far[j+1]
			}
		}
		j = j - 1
	}
	for idx := 0; idx < len(prices)-1; idx++ {
		curr_profit := max_so_far[idx+1] - prices[idx]
		if curr_profit > max_profit {
			max_profit = curr_profit
		}
	}
	return max_profit
}