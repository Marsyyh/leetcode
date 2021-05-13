/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
/*

 */
func isValidSudoku(board [][]byte) bool {
	// iterate once, check all line, row, square
	line, row, sq := NewSudokuMap(), NewSudokuMap(), NewSudokuMap()
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			cur := board[j][i]
			if cur == '.' {
				continue
			}
			if !line.add(i, cur) {
				return false
			}
			if !row.add(j, cur) {
				return false
			}
			if !sq.add(idxSq(i, j), cur) {
				return false
			}
		}
	}

	return true
}

// using bit storage
type SudokuMap []int

func NewSudokuMap() SudokuMap {
	return SudokuMap(make([]int, 9))
}

// find idx storage | num
func (m SudokuMap) add(idx int, num byte) (ok bool) {
	n := int(num - '0')
	cur := m[idx]
	cur = cur | (1 << (n - 1))
	// if not same, num not exist in storage
	if cur != m[idx] {
		m[idx] = cur
		ok = true
		return
	}
	ok = false
	return
}

func idxSq(line, row int) int {
	return line/3 + (row/3)*3
}

// @lc code=end

