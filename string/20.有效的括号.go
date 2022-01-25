package string

func isValid(s string) bool {
	length := len(s)

	if length%2 != 0 {
		return false
	}

	hash := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack []byte

	for i := 0; i < length; i++ {
		if hash[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != hash[s[i]] {
				return false
			}

			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
