package array

func maxSubArray(nums []int) int {
	max := nums[0]
	sum := 0

	i := 0
	for i < len(nums) {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}

		if sum > max {
			max = sum
		}

		i++
	}

	return max
}
