/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
/*
	abba
	dp[i][j] represent from  i to j  is palindrome
	init d[i][i] = 1, self to self is palindrome
	1  0  0  1
	   1  1  0
		  1  0
		     1

*/
func longestPalindrome(s string) string {
	// init
	dp := make([][]bool, len(s))
	for j := 0; j < len(dp); j++ {
		dp[j] = make([]bool, len(s))
	}

	r := []rune(s)
	// iterate from bot to top
	lgstLeft := 0
	lgstRight := 0
	for i := 0; i < len(dp); i++ {
		for j := i; j >= 0; j++ {
			if r[i] == r[j] && (i == j || j == i-1 || dp[j-1][i-1]) {
				dp[j][i] = true
				if i-j > lgstRight-lgstLeft {
					lgstLeft = i
					lgstRight = j
				}
			}
		}
	}
	return string(rune[lgstLeft : lgstRight+1])
}

// @lc code=end

