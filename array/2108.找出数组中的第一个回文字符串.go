package array

func firstPalindrome(words []string) string {
	for _, word := range words {
		isPalindrome := true

		for i := 0; i < len(word)/2; i++ {
			if word[i] != word[len(word)-1-i] {
				isPalindrome = false
				break
			}
		}
		if isPalindrome {
			return word
		}
	}

	return ""
}

//func firstPalindrome(words []string) string {
//	for _, word := range words {
//		if len(word) == 1 {
//			return word
//		}
//
//		i := 0
//		j := len(word) - 1
//		isSame := true
//
//		for i < j {
//			if word[i] != word[j] {
//				isSame = false
//				break
//			}
//			i++
//			j--
//		}
//
//		if isSame {
//			return word
//		}
//	}
//	return ""
//}
