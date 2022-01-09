package array

import "sort"

func majorityElement(nums []int) int {
	sort.Ints(nums)

	return nums[len(nums)/2]
}

//func majorityElement(nums []int) int {
//	count := math.Trunc(float64(len(nums)) / 2)
//
//	m := make(map[int]int)
//	for _, v := range nums {
//		m[v]++
//
//		if m[v] > int(count) {
//			return v
//		}
//	}
//
//	return -1
//}
