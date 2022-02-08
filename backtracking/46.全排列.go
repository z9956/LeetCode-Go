package backtracking

func permute(nums []int) [][]int {
	var res [][]int
	var dfs func([]int)

	dfs = func(path []int) {
		if len(nums) == 0 {
			p := make([]int, len(path))
			copy(p, path)
			res = append(res, p)
		}
		for i := 0; i < len(nums); i++ {
			cur := nums[i]
			path = append(path, cur)
			nums = append(nums[:i], nums[i+1:]...)
			dfs(path)
			nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
			path = path[:len(path)-1]

		}
	}

	dfs([]int{})
	return res
}
