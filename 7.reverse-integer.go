/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
	var res int32
	const MaxUint = ^uint32(0)
	const MaxInt = int32(MaxUint >> 1)
	for x != 0 {
		if Abs(res) > MaxInt / 10 {
			return 0
		}
		res *= 10
		res += int32(x) % 10
		x = x / 10
	}
	return int(res)
}

func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

