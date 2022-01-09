package array

func merge(nums1 []int, m int, nums2 []int, n int) {
	left := m - 1
	right := n - 1
	index := m + n - 1

	for left >= 0 || right >= 0 {
		if left < 0 {
			nums1[index] = nums2[right]
			right--
		} else if right < 0 {
			nums1[index] = nums1[left]
			left--
		} else if nums1[left] > nums2[right] {
			nums1[index] = nums1[left]
			left--
		} else {
			nums1[index] = nums2[right]
			right--
		}
		index--
	}
}
