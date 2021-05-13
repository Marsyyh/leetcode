/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if len(haystack) == 0 && len(needle) == 0 {
		return 0
	}

	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		// check every start with needle
		if len(needle) > len(haystack)-i {
			return -1
		}
		find := true
		for j, k := 0, i; k < len(haystack) && j < len(needle); j, k = j+1, k+1 {
			if haystack[k] != needle[j] {
				find = false
				break
			}
		}
		if find {
			return i
		}
	}
	return -1
}

// @lc code=end

