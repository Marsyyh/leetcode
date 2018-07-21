package main

import "fmt"

func main() {
	test := []int{1, 1, 2}
	res := permuteUnique(test)
	fmt.Println(res)
}

// dfs
/*
	             1123
	        /              \            \
	1(123) 11swap      2(113) 13 swap    3(112) 14 swap   12 don't swap because  1 exist and swap once
2, 3 swap, 2,4 swap
Time: O(n!) worst case, for step 1 n, step 2 n*n-1, .... n!
Space: O(nlogn), every step worst case n for map, logn step
*/
func permuteUnique(nums []int) [][]int {
	re := &[][]int{}
	if nums == nil || len(nums) == 0 {
		return *re
	}
	helper(&nums, 0, re)
	return *re
}

func helper(nums *[]int, layer int, re *[][]int) {
	if layer == len(*nums) {
		tmp := make([]int, len(*nums))
		copy(tmp, *nums)
		*re = append(*re, tmp)
		return
	}

	m := make(map[int]int)
	// inner loop for swap
	for i := layer; i < len(*nums); i++ {
		// if in map, continue
		if _, ok := m[(*nums)[i]]; ok {
			continue
		}
		// otherwise add to map and swap and go deeper
		m[(*nums)[i]] = i
		(*nums)[i], (*nums)[layer] = (*nums)[layer], (*nums)[i]
		helper(nums, layer+1, re)
		// back trace, switch back
		(*nums)[i], (*nums)[layer] = (*nums)[layer], (*nums)[i]
	}
}
