package hash_Table

func canBeEqual(target []int, arr []int) bool {
	length := len(arr)

	if length == 1 {
		return target[0] == arr[0]
	}

	m := make(map[int]int)

	for _, num := range target {
		m[num]++
	}

	for _, num := range arr {
		if m[num] == 0 {
			return false
		}
		m[num]--
	}

	return true
}
