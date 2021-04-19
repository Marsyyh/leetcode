/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	re := []byte{}
	i := 0
	for true {
		if i >= len(strs[0]) {
			return string(re)
		}

		cur := strs[0][i]
		for _, s := range strs {
			if i >= len(s) || s[i] != cur {
				return string(re)
			}
		}
		re = append(re, cur)
		i++
	}

	return string(re)
}

// @lc code=end

