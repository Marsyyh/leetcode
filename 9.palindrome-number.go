/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x != intReverse(x) {
		return false
	}
	return true
}

func intReverse(x int) int {
	res := 0
	for x != 0 {
		res *= 10
		res += x % 10
		x = x / 10
	}
	return res
}

// @lc code=end

