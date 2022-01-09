package array

func singleNumber(nums []int) int {
	ans := 0

	for i := 0; i < len(nums); i++ {
		ans ^= nums[i]
	}

	return ans
}
