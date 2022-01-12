package array

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var word1Str string
	var word2Str string

	for i := 0; i < len(word1); i++ {
		word1Str += word1[i]
	}

	for i := 0; i < len(word2); i++ {
		word2Str += word2[i]
	}

	return word1Str == word2Str
}
