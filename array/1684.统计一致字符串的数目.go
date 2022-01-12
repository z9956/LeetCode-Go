package array

func countConsistentStrings(allowed string, words []string) int {
	count := 0
	m := make(map[int32]bool)

	for _, char := range allowed {
		m[char] = true
	}

	for _, chars := range words {
		status := true

		for _, word := range chars {
			if !m[word] {
				status = false
				break
			}
		}

		if !status {
			continue
		}
		count++
	}

	return count
}
