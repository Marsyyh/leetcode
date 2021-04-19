/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	var stack []byte
	b := []byte(s)
	for _, elem := range b {
		switch elem {
		case '{':
			stack = append(stack, '{')
		case '(':
			stack = append(stack, '(')
		case '[':
			stack = append(stack, '[')
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

// @lc code=end

