/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 */

// @lc code=start
func countPrimes(n int) int {
	m := make(map[int]bool)
	count := 0
	for i := 2; i < n; i++ {
		if !m[i] {
			count++
			for j := 2; j * i < n; j++ {
				m[j*i] = true
			}
		}
	}
	return count
}
// @lc code=end

