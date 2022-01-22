package hash_Table

func intersection(nums1 []int, nums2 []int) []int {
	var results []int
	m := make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] = 1
	}

	for i := 0; i < len(nums2); i++ {
		if m[nums2[i]] > 0 {
			m[nums2[i]]++
		}

		if m[nums2[i]] == 2 {
			results = append(results, nums2[i])
		}
	}

	return results
}
