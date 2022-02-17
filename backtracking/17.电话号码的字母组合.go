package backtracking

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var results []string
	var dfs func(string, int)
	hash := map[string]string{
		"0": "",
		"1": "",
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	dfs = func(str string, i int) {
		if i == len(digits) {
			results = append(results, str)
			return
		}
		for _, c := range hash[string(digits[i])] {
			dfs(str+string(c), i+1)
		}
	}

	dfs("", 0)

	return results
}
