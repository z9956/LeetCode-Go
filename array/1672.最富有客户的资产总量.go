package array

func maximumWealth(accounts [][]int) int {
	var max int

	for _, account := range accounts {
		sum := 0

		for _, value := range account {
			sum += value
		}

		if sum > max {
			max = sum
		}
	}

	return max
}
