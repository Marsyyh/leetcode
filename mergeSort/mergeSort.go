package mergeSort

func mergeSort(input []int) []int {
	helpArr := make([]int, len(input))
	sort(&input, &helpArr, 0, len(input)-1)
	return input
}

func sort(input, helpArr *[]int, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	sort(input, helpArr, left, mid)
	sort(input, helpArr, mid+1, right)
	merge(input, helpArr, left, right, mid)
}

func merge(input, helper *[]int, left, right, mid int) {
	var i int
	j := mid + 1
	for i = left; i <= right; i++ {
		(*helper)[i] = (*input)[i]
	}
	i = left
	for i <= mid && j <= right {
		(*input)[left], i, j = getValue(helper, i, j)
		left++
	}
	for i <= mid {
		(*input)[left] = (*helper)[i]
		left++
		i++
	}
	for j <= right {
		(*input)[left] = (*helper)[j]
		left++
		j++
	}
}

func getValue(input *[]int, i, j int) (int, int, int) {
	if (*input)[i] < (*input)[j] {
		return (*input)[i], i + 1, j
	}
	return (*input)[j], i, j + 1
}
