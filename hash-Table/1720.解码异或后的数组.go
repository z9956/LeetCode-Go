package hash_Table

func decode(encoded []int, first int) []int {
	length := len(encoded) + 1
	decoded := make([]int, length)
	decoded[0] = first

	for i := 1; i < length; i++ {
		decoded[i] = decoded[i-1] ^ encoded[i-1]
	}

	return decoded
}
