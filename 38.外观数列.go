/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start
func countAndSay(n int) string {
	prev := "1"
	for i := 2; i <= n; i++ {
		cur := ""
		for j := 0; j < len(prev); j++ {
			// find same byte
			nn := 1
			for j < len(prev)-1 && prev[j] == prev[j+1] {
				j++
				nn++
			}
			cur = cur + fmt.Sprintf("%d%s", nn, string(prev[j]))
		}
		prev = cur
	}

	return prev
}

// @lc code=end

