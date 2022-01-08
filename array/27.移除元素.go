package array

func removeElement(nums []int, val int) int {
	i := 0

	for i < len(nums) {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	return i
}
