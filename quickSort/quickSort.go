package quickSort

func quickSort(input []int) []int {
	helper(input, 0, len(input)-1)
	return input
}

func helper(input []int, left, right int) {
	if left >= right {
		return
	}
	idx := partition(left, right)
	for i := left; i < idx; i++ {
		if input[i] >= input[idx] {
			input[i], input[idx] = input[idx], input[i]
			idx--
			input[i], input[idx] = input[idx], input[i]
			i--
		}
	}
	helper(input, left, idx-1)
	helper(input, idx+1, right)
}

// Place to get the partition logic
// In this we simply use the most right
func partition(left, right int) int {
	return right
}
