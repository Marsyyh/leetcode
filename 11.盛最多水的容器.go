/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
/*
solution 1: dp, Speed: O(n^2), Space: O(n^2), Space Complexity exceed
dp[i, j] means maxArea from i to j
0   a(1,0)  a(2,0) = max(a(2,0),a(1,0),a(2,0))
    0		a(2,1)
				0	a(3,2)
					0

Solution 2:
i, j double pointer i: left most, j: right most

moving shorter side, may increse the maxmum
*/
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := calArea(height[i], height[j], j-i)
	for i < j {
		// moving i
		if height[i] < height[j] {
			i++
			max = maxInt(max, calArea(height[i], height[j], j-i))
			continue
		}
		j--
		max = maxInt(max, calArea(height[i], height[j], j-i))
		// moving j
	}

	return max
}

func maxAreaDP(height []int) int {
	n := len(height)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	max := 0
	for i := 1; i < len(height); i++ {
		for j := i - 1; j >= 0; j-- {
			area := calArea(height[i], height[j], i-j)
			if j+1 == i {
				dp[i][j] = area
				max = maxInt(max, area)
				continue
			}
			m := maxThreeInt(area, dp[i-1][j], dp[i][j+1])
			dp[i][j] = m
			max = maxInt(max, m)
		}
	}
	return max
}

func calArea(a, b, l int) int {
	if a > b {
		return b * l
	}
	return a * l
}

func maxThreeInt(a, b, c int) int {
	m := maxInt(a, b)
	return maxInt(m, c)
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

