/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
/*
	dfs
*/
func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(candidates)
	dfs(candidates, 0, target, &[]int{}, &result)
	return result
}

func dfs(cans []int, idx, target int, cur *[]int, result *[][]int) {
	if target == 0 {
		now := make([]int, len(*cur))
		copy(now, *cur)
		*result = append(*result, now)
		return
	}

	if target < 0 || idx >= len(cans) {
		return
	}

	// fill & go to next
	// for same item, prune
	n, now, curIdx := 1, cans[idx], len(*cur)
	for idx < len(cans)-1 && cans[idx] == cans[idx+1] {
		idx++
		n++
	}

	// not fill & go to next
	dfs(cans, idx+1, target, cur, result)

	for i := 1; i <= n; i++ {
		*cur = append(*cur, now)
		dfs(cans, idx+1, target-now*i, cur, result)
	}

	*cur = (*cur)[:curIdx]
}

// @lc code=end

