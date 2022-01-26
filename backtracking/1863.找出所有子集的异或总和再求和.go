package backtracking

func subsetXORSum(nums []int) int {
	var sum int
	var dfs func(index int, val int)

	dfs = func(index int, val int) {
		if index == len(nums) {
			sum += val
			return
		}

		dfs(index+1, val^nums[index])
		dfs(index+1, val)
	}

	dfs(0, 0)

	return sum
}
