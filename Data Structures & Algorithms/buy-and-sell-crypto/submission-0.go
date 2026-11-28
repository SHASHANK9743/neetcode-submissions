func maxProfit(prices []int) int {
	max_profit := 0
	for idx := 0; idx < len(prices)-1; idx++ {
		for idx1 := idx + 1; idx1 < len(prices); idx1++ {
			curr_profit := prices[idx1] - prices[idx]
			if curr_profit > max_profit {
				max_profit = curr_profit
			}
		}
	}
	return max_profit
}