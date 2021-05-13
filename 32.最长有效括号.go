/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
/*
	dp:
	dp[i] current index longest valid length
	if s[i] == '(', dp[i] = 0
	if s[i] == ')', dp[i] = max(
	case1	s[i-1] == '(' ?
				dp[i-2] + 2,
	case2	s[i-1] == ')' && s[i-dp[i-1]-1] == '(' ?
				dp[i-1] + 2

		connect with before, only happens in case2
			dp[i] = dp[i] + dp[i-dp[i]]
	)
	else 0
*/
func longestValidParentheses(s string) int {
	max := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		// open, 0
		if s[i] == '(' {
			dp[i] = 0
		}

		// close
		if s[i] == ')' {
			// case 1, s[i-1] == '('
			if s[i-1] == '(' {
				if i > 1 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else {
				// case 2
				if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
					dp[i] = dp[i-1] + 2
				}

				if i-dp[i] > 0 {
					dp[i] += dp[i-dp[i]]
				}
			}
			if max < dp[i] {
				max = dp[i]
			}
		}
	}
	return max
}

// @lc code=end

