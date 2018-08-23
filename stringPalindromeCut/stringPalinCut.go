package stringPalindromeCut

func MinCuts(s string) int {
	dp := make([]int, len(s)+1)
	fillDp(s, dp)
	return dp[len(s)]
}

func fillDp(s string, dp []int) {
	// make dp[0] == -1 because it can easy all string + 1 = 0
	dp[0] = -1
	dp[1] = 0
	rs := []rune(s)
	// isPalin[i][j] is from index i to index j, isPalin
	isPalin := initialIsPalin(len(s))

	for i := 2; i <= len(rs); i++ {
		calLayer(i, rs, dp, isPalin)
	}
}

func calLayer(n int, rs []rune, dp []int, isPalin [][]bool) {
	dp[n] = 1<<31 - 1
	for i := 0; i < n; i++ {
		if isPalin[i+1][n-1] && rs[i] == rs[n-1] {
			isPalin[i][n] = true
			if dp[n] > dp[i]+1 {
				dp[n] = dp[i] + 1
			}
		}
	}
}

func initialIsPalin(s int) [][]bool {
	re := make([][]bool, s+1)
	for i := range re {
		re[i] = make([]bool, s+1)
	}
	for i := 0; i < s; i++ {
		re[i][i] = true
		re[i+1][i] = true
		re[i][i+1] = true
	}
	re[s][s] = true
	return re
}
