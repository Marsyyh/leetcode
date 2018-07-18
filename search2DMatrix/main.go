package main

import "fmt"

func main() {
	test := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target := 5
	res := searchMatrix(test, target)
	fmt.Print(res)
}

// for int[m][n], it's 1d index = len(int[0]) * n + m
// for index, m = index % len(int[0])
// 			  n = index / len(int[0])
func searchMatrix(matrix [][]int, target int) bool {
	result := false
	return result
}
