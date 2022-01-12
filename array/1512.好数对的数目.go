package array

func numIdenticalPairs(nums []int) int {
	var count int
	var length = len(nums)

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if nums[i] == nums[j] && i < j {
				count++
			}
		}
	}

	return count
}
