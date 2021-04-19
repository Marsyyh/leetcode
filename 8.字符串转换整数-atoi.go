/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	const (
		MaxInt32 = 1<<31 - 1
		MinInt32 = -1 << 31
	)

	rn := []rune(s)
	result := 0
	neg := false
	seen := false
	for _, r := range rn {
		// filter prefix
		if !seen && r == ' ' {
			continue
		}

		if !seen && r == '+' {
			seen = true
			continue
		}

		// neg positive
		if !seen && r == '-' {
			neg = !neg
			seen = true
			continue
		}

		d := r - '0'
		if !seen && d >= 0 && d < 10 {
			seen = true
			result = int(d)
			continue
		}

		if d > 9 || d < 0 {
			return genResult(result, neg)
		}

		result = result*10 + int(d)

		if !neg && result > MaxInt32 {
			return MaxInt32
		}

		if neg && -result < MinInt32 {
			return MinInt32
		}
	}

	return genResult(result, neg)
}

func genResult(r int, neg bool) int {
	if neg {
		return -r
	}
	return r
}

// @lc code=end

