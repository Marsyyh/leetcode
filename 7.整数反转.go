/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	const MaxInt32 = 1<<31 - 1
	const MinInt32 = -1 << 31

	neg := false
	if x < 0 {
		neg = true
		x = -x
	}

	idx := 0
	result := 0
	for x > 0 {
		bit := x % 10
		x = x / 10
		result = result*10 + bit

		// exceed max or min
		if (result > MaxInt32 && !neg) || (-result < MinInt32 && neg) {
			return 0
		}
		idx++
	}

	if neg {
		return -result
	}
	return result
}

// @lc code=end

