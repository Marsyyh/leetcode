package wildcardMatching

/*
	   a a b b b c a a
	a  1 0 0 0 0 0 0 0
	*  1 1 1 1 1 1 1 1
	b  * 0 1 1 1 0 0 0
	*  * 0 1 1 1 1 1 1
	*  * * 1 1 1 1 1 1
	?  * * * 1 1 1 1 1
	c  * * * * 0 1 0 0
	a  * * * * * 0 1 0

*/
func wildcard(input string, pattern string) bool {
	dp := make([]bool, len(input))
	cur := make([]bool, len(input))
	curIndex := 0
	for i := 0; i < len(pattern); i++ {
		for j := curIndex; j < len(input); j++ {
			prev := true
			cur[j] = false
			if j > 0 {
				prev = dp[j-1]
			}
			if pattern[i] == '*' {
				if curIndex == 0 {
					cur[j] = true
				} else {
					cur[j] = dp[j] || cur[j-1]
				}
			} else if pattern[i] == '?' || pattern[i] == input[j] {
				cur[j] = prev
			}
		}
		if pattern[i] != '*' {
			curIndex++
		}
		copy(dp, cur)
	}
	return dp[len(input)-1]
}
