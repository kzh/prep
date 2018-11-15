package arrays

// Even Odd
func EvenOdd(arr []int) []int {
	clone := make([]int, len(arr))
	copy(clone, arr)

	front, back := 0, len(clone)-1
	for front < back {
		if clone[front]%2 == 0 {
			front++
		} else {
			temp := clone[front]
			clone[front] = clone[back]
			clone[back] = temp
			back--
		}
	}

	return clone
}
