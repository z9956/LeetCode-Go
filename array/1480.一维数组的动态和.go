package array

func runningSum(nums []int) []int {
	var res []int
	sum := 0

	for _, num := range nums {
		sum += num
		res = append(res, sum)
	}

	return res
}
