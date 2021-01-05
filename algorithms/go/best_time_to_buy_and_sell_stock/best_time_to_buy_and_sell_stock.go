package best_time_to_buy_and_sell_stock

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
			continue
		}

		current := price - minPrice
		if current > maxProfit {
			maxProfit = current
		}
	}

	return maxProfit
}
