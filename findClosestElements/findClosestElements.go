package findClosestElements

import (
	"sort"
)

// binary search find most close number
// then i, j pointer move smallest abs
func findClosestElements(arr []int, k int, x int) []int {
	var left, right int
	result := make([]int, k)
	// find most left closest
	for left, right = 0, len(arr); left < right; {
		mid := left + (right-left)/2
		if x > arr[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left > 0 {
		left--
	}
	// now x in the [left, left + 1]
	for i, j := left, left+1; k > 0; k-- {
		if i < 0 || j < len(arr) && x-arr[i] > arr[j]-x {
			result[k-1] = arr[j]
			j++
		} else {
			result[k-1] = arr[i]
			i--
		}
	}
	// we can use two array each one put small each one put larger and combine them together
	sort.Ints(result)
	return result
}
