package backtracking

func subsets(nums []int) [][]int {
	var res [][]int
	var tmp []int
	var dfs func(int, []int)

	dfs = func(i int, tmp []int) {
		res = append(res, append([]int{}, tmp...))
		for j := i; j < len(nums); j++ {
			tmp = append(tmp, nums[j])
			dfs(j+1, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(0, tmp)
	return res
}
