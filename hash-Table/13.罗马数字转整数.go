package hash_Table

func romanToInt(s string) (sum int) {
	table := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := range s {
		cur := table[s[i]]

		if i < len(s)-1 && cur < table[s[i+1]] {
			sum -= cur
		} else {
			sum += cur
		}
	}

	return sum
}
