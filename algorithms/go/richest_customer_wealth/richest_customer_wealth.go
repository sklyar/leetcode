package richest_customer_wealth

func maximumWealth(accounts [][]int) int {
	var max int
	for _, customer := range accounts {
		var wealth int
		for _, account := range customer {
			wealth += account
		}

		if wealth > max {
			max = wealth
		}
	}

	return max
}
