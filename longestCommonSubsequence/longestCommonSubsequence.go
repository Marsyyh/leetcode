package longestCommonSubsequence

func Lcs(source string, target string) int {
	// corner
	if len(source) == 0 || len(target) == 0 {
		return 0
	}
	// initial
	dp := make([]int, len(source))
	for i := range source {
		if source[i] == target[0] {
			dp[i] = 1
		} else if i > 0 {
			dp[i] = dp[i-1]
		}
	}
	// dp loop
	max := 0
	for j := 1; j < len(target); j++ {
		cur := make([]int, len(dp))
		copy(cur, dp)
		for i := range source {
			plus := 0
			if source[i] == target[j] {
				plus = 1
			}
			if i == 0 {
				if plus == 1 {
					dp[0] = 1
				}
			} else {
				if plus == 1 {
					dp[i] = cur[i-1] + 1
				} else {
					if dp[i-1] > cur[i] {
						dp[i] = dp[i-1]
					} else {
						dp[i] = cur[i]
					}
				}
			}
			if max < dp[i] {
				max = dp[i]
			}
		}
	}
	return max
}
