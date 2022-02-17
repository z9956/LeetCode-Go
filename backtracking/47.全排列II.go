package backtracking

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int
	var path []int
	var dfs func(map[int]bool)

	dfs = func(vis map[int]bool) {
		if len(path) == len(nums) {
			p := make([]int, len(path))
			copy(p, path)
			res = append(res, p)
		}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && !vis[i-1] {
				continue
			}

			if !vis[i] {
				vis[i] = true
				path = append(path, nums[i])
				dfs(vis)
				path = path[:len(path)-1]
				vis[i] = false
			}

		}
	}

	dfs(make(map[int]bool))
	return res
}
