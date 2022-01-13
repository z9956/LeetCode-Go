package array

func countKDifference(nums []int, k int) int {
	var count int = 0
	length := len(nums)

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if i != j && (nums[i]-nums[j] == k) {
				count++
			}
		}
	}

	return count
}
