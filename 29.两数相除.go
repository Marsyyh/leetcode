/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
// using minus time exceeded
// thinking using bit operator for dividing
// move and minus
/* 		7 / 2
0111 / 0010
0001 >= 0010 ? no result 0, reminder 01
	move reminder & result left 1
0011 >= 0010 ? yes result 1, reminder 1
	move reminder & result left 1
0011 >= 0010 ? yes result 11, reminder 1
	no next, result 11, reminder 1
*/
func divide(dividend, divisor int) int {
	MaxInt := 1<<31 - 1
	MinInt := -MaxInt - 1

	positive := 1
	if dividend < 0 {
		dividend = -dividend
		positive = -1 * positive
	}

	if divisor < 0 {
		divisor = -divisor
		positive = -1 * positive
	}

	q := 0 // quotient
	// 32 bits at most 32 times
	for i := 31; i >= 0; i-- {
		if (dividend >> i) >= divisor {
			q = q + (1 << i)
			dividend = dividend - (divisor << i)
		}
	}
	if q*positive > MaxInt {
		return MaxInt
	}
	if q*positive < MinInt {
		return MinInt
	}
	return q * positive
}

func minusdivide(dividend int, divisor int) int {
	MaxInt := 1<<31 - 1
	MinInt := -MaxInt - 1
	positive := 1
	if dividend < 0 {
		dividend = -dividend
		positive = -1 * positive
	}

	if divisor < 0 {
		divisor = -divisor
		positive = -1 * positive
	}

	n := 0
	for dividend >= divisor {
		dividend = dividend - divisor
		n++
	}
	if n*positive > MaxInt {
		return MaxInt
	}
	if n*positive < MinInt {
		return MinInt
	}
	return n * positive
}

// @lc code=end

