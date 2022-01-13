package array

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	max := 0

	for left < right {
		area := (right - left) * min(height[left], height[right])
		if area > max {
			max = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return max
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}
