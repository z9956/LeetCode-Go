package array

func maxProfit(prices []int) int {
	var max int
	var minPrice int = prices[0]

	for i := 0; i < len(prices); i++ {

		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		var profit = prices[i] - minPrice

		if profit > max {
			max = profit
		}
	}

	return max
}
