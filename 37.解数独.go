/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
/*
	dfs find all
*/
func solveSudoku(board [][]byte) [][]byte {
	sudo := &sudokuSolver{
		lineM: numMap(make([]int, 9)),
		rowM:  numMap(make([]int, 9)),
		sqM:   numMap(make([]int, 9)),
		board: board,
	}
	sudo.Proc()
	return sudo.board
}

type sudokuSolver struct {
	lineM numMap
	rowM  numMap
	sqM   numMap
	board [][]byte
}

func (s *sudokuSolver) Proc() {
	s.init()
	s.iterate(0, 0)
}

// init set all maps
func (s *sudokuSolver) init() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cur := s.board[i][j]
			if cur != '.' {
				s.lineM.add(i, cur)
				s.rowM.add(j, cur)
				s.sqM.add(idxSq(i, j), cur)
			}
		}
	}
}

func (s *sudokuSolver) iterate(i, j int) bool {
	// edge
	if j >= 9 && i < 9 {
		i, j = i+1, 0
		// switch line
		return s.iterate(i, j)
	}

	// end means all good, true
	if i >= 9 {
		return true
	}

	if s.board[i][j] != '.' {
		return s.iterate(i, j+1)
	}

	// fill current, and iterate next
	for k := 1; k < 10; k++ {
		// prune
		if s.add(k, i, j) {
			if s.iterate(i, j+1) {
				// find right answer
				return true
			} else {
				// clean current & for next
				s.remove(k, i, j)
			}
		}
	}

	// if all false, return false
	return false
}

func (s *sudokuSolver) check(num, i, j int) bool {
	if !s.lineM.check(i, num) || !s.rowM.check(j, num) || !s.sqM.check(idxSq(i, j), num) {
		return false
	}
	return true
}

func (s *sudokuSolver) add(num, i, j int) bool {
	if !s.lineM.addInt(i, num) {
		return false
	}

	if !s.rowM.addInt(j, num) {
		// rollback lineM
		s.lineM.remove(i, num)
		return false
	}

	if !s.sqM.addInt(idxSq(i, j), num) {
		// rollback lineM & rowM
		s.lineM.remove(i, num)
		s.rowM.remove(j, num)
		return false
	}

	s.board[i][j] = byte('0' + num)
	return true
}

func (s *sudokuSolver) remove(num, i, j int) {
	s.lineM.remove(i, num)
	s.rowM.remove(j, num)
	s.sqM.remove(idxSq(i, j), num)
	s.board[i][j] = '.'
}

type numMap []int

func (m numMap) dryAdd(idx int, num byte) int {
	n := int(num - '0')
	return m.dryAddInt(idx, n)
}

func (m numMap) dryAddInt(idx, num int) int {
	cur := m[idx]
	return cur | (1 << (num - 1))
}

func (m numMap) check(idx, num int) bool {
	dry := m.dryAddInt(idx, num)
	if dry != m[idx] {
		return true
	}
	return false
}

func (m numMap) add(idx int, num byte) (ok bool) {
	dry := m.dryAdd(idx, num)
	// if not same, num not exist in storage
	if dry != m[idx] {
		m[idx] = dry
		ok = true
		return
	}
	ok = false
	return
}

func (m numMap) addInt(idx, num int) (ok bool) {
	dry := m.dryAddInt(idx, num)
	if dry != m[idx] {
		m[idx] = dry
		ok = true
		return
	}
	ok = false
	return
}

func (m numMap) remove(idx, num int) {
	m[idx] = m[idx] ^ (1 << (num - 1))
}

func idxSq(i, j int) int {
	return i/3 + (j/3)*3
}
