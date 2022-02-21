package backtracking

func generateParenthesis(n int) []string {
	var results []string
	var dfs func(int, int, string)

	dfs = func(left int, right int, str string) {
		if len(str) == 2*n {
			results = append(results, str)
			return
		}

		if left > 0 {
			dfs(left-1, right, str+"(")
		}

		if left < right {
			dfs(left, right-1, str+")")
		}
	}

	dfs(n, n, "")

	return results
}
