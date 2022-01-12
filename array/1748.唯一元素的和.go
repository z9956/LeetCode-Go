package array

func sumOfUnique(nums []int) int {
	var sum int
	repeated := make(map[int]int)

	for _, num := range nums {
		repeated[num]++
	}

	for key, num := range repeated {
		if num == 1 {
			sum += key
		}
	}
	return sum
}
