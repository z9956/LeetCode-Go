package array

func findErrorNums(nums []int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	ans := make([]int, 2)

	for i := 1; i <= len(nums); i++ {
		if m[i] == 2 {
			ans[0] = i
		} else if m[i] == 0 {
			ans[1] = i
		}
	}

	return ans
}
