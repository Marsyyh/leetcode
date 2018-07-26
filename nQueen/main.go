package main

import (
	"bytes"
	"fmt"
)

// Using dfs
/*
	n branches, using slice, index means row, value means the col.
	O(1) check valid:
	[]bool usedCol for the col used size n
	[]bool usedDiag for diag used size 2n - 1
	[]bool usedRevDiag for reverse diag used size 2n - 1
*/
func solveNQueens(n int) [][]string {
	re := make([]string, n)
	usedCol := make([]bool, n)
	usedDiag := make([]bool, 2*n-1)
	usedRevDiag := make([]bool, 2*n-1)
	var result [][]string
	helper(n, 0, re, usedCol, usedDiag, usedRevDiag, &result)
	return result
}

func helper(n int, layer int, re []string, usedCol []bool,
	usedDiag []bool, usedRevDiag []bool, result *[][]string) {
	//basecase
	if layer == n {
		tmp := make([]string, n)
		copy(tmp, re)
		*result = append(*result, tmp)
		return
	}

	for i := 0; i < n; i++ {
		if valid(n, i, layer, usedCol, usedDiag, usedRevDiag) {
			mark(n, i, layer, usedCol, usedDiag, usedRevDiag)
			re[layer] = generateStr(i, n)
			helper(n, layer+1, re, usedCol, usedDiag, usedRevDiag, result)
			re[layer] = ""
			unmark(n, i, layer, usedCol, usedDiag, usedRevDiag)
		}
	}
}

func valid(n int, i int, layer int, usedCol []bool, usedDiag []bool, usedRevDiag []bool) bool {
	return !usedCol[i] && !usedDiag[layer-i+n-1] && !usedRevDiag[layer+i]
}

func mark(n int, i int, layer int, usedCol []bool, usedDiag []bool, usedRevDiag []bool) {
	usedCol[i] = true
	/*	2 [0,0][1,1][2,2] 3[1,0][2,1] 4[2,0]
		1 [0,1][1,2]
		0[0,2]
		col - row + n - 1
	*/
	usedDiag[layer-i+n-1] = true
	/*
		4	3	2[0,2][1,1][2,0]
				1[1,2][2,1]
				0[2,2]
	*/
	usedRevDiag[layer+i] = true
}

func unmark(n int, i int, layer int, usedCol []bool, usedDiag []bool, usedRevDiag []bool) {
	usedCol[i] = false
	usedDiag[layer-i+n-1] = false
	usedRevDiag[layer+i] = false
}

func generateStr(i int, size int) string {
	var buffer bytes.Buffer
	for j := 0; j < size; {
		if j == i {
			buffer.WriteString("Q")
		} else {
			buffer.WriteString(".")
		}
		j++
	}
	return buffer.String()
}

func main() {
	res := solveNQueens(5)
	fmt.Println(res)
}
