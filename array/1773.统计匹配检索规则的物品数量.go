package array

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var count int

	var keys = map[string]int{
		"type":  0,
		"color": 1,
		"name":  2,
	}

	for _, item := range items {
		if item[keys[ruleKey]] == ruleValue {
			count++
		}
	}

	return count
}
