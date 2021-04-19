/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */

// @lc code=start
func isMatch(s string, p string) bool {
	// cornercase
    if !validatePattern(p) {
		return false
	}
	if len(s) == 0 {
		if len(p) == 0 {
			return true
		}
		return false
	}

	i, j := 0, 0
	sr := []rune(s)
	pr := []rune(p)
	matchPat := match(pr[0])
	meet := false
	for i < len(sr) && j < len(pr) {
		cur := sr[i]
		pat := pr[j]
		if matchPat(cur) {
			i++
			if pat == '*' {
				meet = true
			} else {
				matchPat = match(pat)
				j++
			}
		} else {
			if !meet {
				return false
			}
			meet = false
			j++
		}
	}
	if i < len(s) {
		return false
	}
	if j == len(s) - 1 && []rune(s)[j] == '*' {
		return true
	}
	return false
}

func match(pattern rune) func(rune) bool {
	return func (m rune) bool {
		if pattern == '.' {
			return true
		}
		if pattern == m {
			return true
		}
		return false
	}
}


func validatePattern(p string) bool {
	if []rune(p)[0] == '*' {
		return false
	}
	return true
}
// @lc code=end

