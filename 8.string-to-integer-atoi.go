/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start
const MAX_UINT = 1 << 31 - 1
const MAX_NINT = - MAX_UINT - 1

func myAtoi(str string) int {
	minus := false
	i, res := 0, 0
	b := []byte(str)

	// trim space
	for _, c := range(b) {
		if c == ' ' {
			i++
		} else {
			break
		}
	}
	b = b[i:]

	// negitive
	if len(b) > 1 && b[0] == '-' {
		minus = true
		b = b[1:]
	} else if len(b) > 1 && b[0] == '+' {
		b = b[1:]
	}

	for j := 0; j < len(b); j++ {
		if b[j] < '0' || b[j] > '9' {
			return Result(res, minus)
		}
		res = res * 10 + int(b[j] - '0')
		if res > MAX_UINT {
			return Result(res, minus)
		}
	}
	return Result(res, minus)
}

func Result(res int, minus bool) int {

	if minus {
		res = -res
		if res < MAX_NINT {
			return MAX_NINT
		}
		return res
	}

	if res > MAX_UINT {
		return MAX_UINT
	}
	return res
}
// @lc code=end

