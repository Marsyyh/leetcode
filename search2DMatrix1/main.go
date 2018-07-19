package main

import "fmt"

func main() {
	test := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	target := 2
	res := searchMatrix(test, target)
	fmt.Print(res)
}

// for int[m][n], it's 1d index = len(int[0]) * n + m
// for index, m = index % len(int[0])
// 			  n = index / len(int[0])
func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	for left, right := 0, len(matrix[0])*len(matrix)-1; left <= right; {
		mid := left + (right-left)/2
		cur := matrix[mid/len(matrix[0])][mid%len(matrix[0])]
		if cur == target {
			return true
		} else if cur > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
