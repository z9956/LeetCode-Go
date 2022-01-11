package array

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var ans []int

	for _, num := range nums1 {
		for i := 0; i < len(nums2); i++ {
			if num == nums2[i] {
				if i == len(nums2)-1 {
					ans = append(ans, -1)
				} else {
					for j := i; j < len(nums2); j++ {
						if nums2[j] > num {
							ans = append(ans, nums2[j])
							break
						} else if j == len(nums2)-1 {
							ans = append(ans, -1)
						}
					}
				}
			}
		}
	}
	return ans
}
