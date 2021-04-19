/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	r := []byte(s)
	// max & sliding window
	// rt: right most + 1
	// lt: left most
	max, lt, rt := 0, 0, 0
	for i, c := range r {
		ci, ok := m[c]
		if ok && ci >= lt {
			fmt.Println(ci, lt, rt, max)
			if rt-lt > max {
				max = rt - lt
			}
			lt = ci + 1
		}
		m[c] = i
		rt++
	}

	if rt-lt > max {
		max = rt - lt
	}

	return max
}

// @lc code=end

