/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	return x == inverseInt(x)
}

func inverseInt(x int) int {
	res := 0
	for x > 0 {
		r := x % 10
		res = res*10 + r
		x = x / 10
	}
	return res
}

// @lc code=end

