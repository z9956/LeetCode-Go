package hash_Table

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}

	m := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if m[nums[i]] {
			return true
		}

		m[nums[i]] = true

		if len(m) > k {
			delete(m, nums[i-k])
		}
	}

	return false
}
