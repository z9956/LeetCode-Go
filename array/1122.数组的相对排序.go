package array

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var res []int
	var tmp []int
	m := make(map[int]bool)

	for _, val := range arr2 {
		m[val] = true

		for j := 0; j < len(arr1); j++ {
			if arr1[j] == val {
				res = append(res, arr1[j])
				arr1 = append(arr1[:j], arr1[j+1:]...)
				j--
			}
		}
	}

	for _, val := range arr1 {
		if m[val] == false {
			tmp = append(tmp, val)
		}
	}

	sort.Ints(tmp)

	return append(res, tmp...)
}
