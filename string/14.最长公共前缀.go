package string

func longestCommonPrefix(strs []string) string {
	hash := make(map[int][]uint8)
	length := len(strs)
	min := min(strs)
	var result string

	if length == 0 || min == 0 {
		return ""
	}

	for i := 0; i < min; i++ {
		for j := 0; j < length; j++ {
			hash[j] = append(hash[j], strs[j][i])
		}
	}

	for i := 0; i < min; i++ {
		for k := 0; k < length; k++ {

			if hash[0][i] != hash[k][i] {
				return result
			} else if k+1 == length {
				result += string(hash[k][i])
			}

		}

	}
	return result
}

func min(list []string) int {
	if len(list) == 0 {
		return 0
	}

	var length int = len(list[0])

	for _, val := range list {
		valLen := len(val)

		if valLen < length {
			length = valLen
		}
	}

	return length
}
