package splitNumMinSqrSum

import (
	"math"
)

func MinSqrSum(num int) int {
	dp := make([]int, num+1)
	fillDp(num, dp, calLayer)
	return dp[num]
}

func fillDp(n int, dp []int, calLayer func(i int, dp []int)) {
	dp[0] = 0
	dp[1] = 1
	i := 2
	for ; i <= n; i++ {
		calLayer(i, dp)
	}
}

func calLayer(i int, dp []int) {
	dp[i] = 1<<31 - 1
	n := int(math.Sqrt(float64(i)))
	j := 1
	for ; j <= n; j++ {
		idx := j * j
		if dp[i] > dp[i-idx]+1 {
			dp[i] = dp[i-idx] + 1
		}
	}
}
