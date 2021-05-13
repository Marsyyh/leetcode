/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	stack := []rune{}
	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			// push
			stack = append(stack, r)
		}

		if r == ')' || r == ']' || r == '}' {
			if len(stack) == 0 {
				return false
			}
			if !match(stack[len(stack)-1], r) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) != 0 {
		return false
	}
	return true
}

func match(left, right rune) bool {
	if left == '(' && right == ')' {
		return true
	}
	if left == '[' && right == ']' {
		return true
	}
	if left == '{' && right == '}' {
		return true
	}
	return false
}

// @lc code=end

