package array

import "strings"

func finalValueAfterOperations(operations []string) int {
	var sum int

	for _, operation := range operations {

		if strings.Count(operation, "++") == 1 {
			sum++
		} else {
			sum--
		}
	}

	return sum
}
