/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
/*
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/
func romanToInt(s string) int {
	reps := []byte("MDCLXVI")
	nums := []int{1000, 500, 100, 50, 10, 5, 1}
	m := make(map[byte]int)
	for i, b := range reps {
		m[b] = nums[i]
	}

	sum := 0
	for i := range s {
		if i != len(s)-1 && m[s[i]] < m[s[i+1]] {
			sum -= m[s[i]]
			continue
		}
		sum += m[s[i]]
	}
	return sum
}

// @lc code=end

