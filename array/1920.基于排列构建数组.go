package array

func buildArray(nums []int) []int {
	var ans []int
	length := len(nums)

	for i := 0; i < length; i++ {
		ans = append(ans, nums[nums[i]])
	}
	return ans
}
