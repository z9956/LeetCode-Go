package array

func smallerNumbersThanCurrent(nums []int) []int {
	var res []int
	var length = len(nums)

	for i := 0; i < length; i++ {
		var count = 0
		for j := 0; j < length; j++ {
			if nums[i] > nums[j] {
				count++
			}
		}
		res = append(res, count)
	}

	return res
}
