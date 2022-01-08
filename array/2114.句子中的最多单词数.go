package array

import "strings"

func mostWordsFound(sentences []string) (max int) {
	for _, sentence := range sentences {
		count := strings.Count(sentence, " ") + 1

		if count > max {
			max = count
		}
	}
	return max
}
