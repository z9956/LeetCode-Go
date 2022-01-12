package array

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)

	for _, v := range arr {
		m[v]++
	}

	res := make(map[int]bool)

	for _, v := range m {
		if res[v] {
			return false
		}
		res[v] = true
	}

	return true
}
