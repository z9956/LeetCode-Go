package array

func searchInsert(nums []int, target int) int {
	length := len(nums)

	left := 0
	right := length - 1

	for left <= right {
		mid := (right-left)>>1 + left

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
