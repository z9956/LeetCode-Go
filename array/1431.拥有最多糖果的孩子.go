package array

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0

	for _, num := range candies {
		if num > max {
			max = num
		}
	}

	res := make([]bool, len(candies))

	for i, num := range candies {
		if num+extraCandies >= max {
			res[i] = true
		}
	}

	return res
}
