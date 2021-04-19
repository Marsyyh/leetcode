/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	container := [][]rune{}
	rs := []rune(s)
	// idx
	i := 0
	for i < len(rs) {
		// fill full row
		fullRow := []rune{}
		for j := 0; j < numRows; j++ {
			if i < len(rs) {
				fullRow = append(fullRow, rs[i])
				i++
			} else {
				fullRow = append(fullRow, rune(0))
			}
		}
		container = append(container, fullRow)

		// fill rest row
		cur := 1

		for cur < numRows-1 {
			if i >= len(rs) {
				break
			}
			row := []rune{}
			for j := 0; j < numRows; j++ {
				if j == numRows-1-cur {
					row = append(row, rs[i])
					i++
				} else {
					row = append(row, rune(0))
				}
			}
			container = append(container, row)
			cur++
		}
	}

	newrs := []rune{}
	for i := 0; i < len(container[0]); i++ {
		for j := 0; j < len(container); j++ {
			if container[j][i] != rune(0) {
				newrs = append(newrs, container[j][i])
			}
		}
	}
	return string(newrs)
}

// @lc code=end

