/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
/*
	dfs

	solusion 1:  multiple branch, const deep
	[2],	 			[2, 2],	 [2, 2, 2]
    [3], [3, 3]

	solution 2: const branch, more deeper
		[2], [3], [6], [7]
*/

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	dfs(candidates, 0, target, &[]int{}, &result)
	return result
}

func dfs(cans []int, idx, target int, cur *[]int, result *[][]int) {
	// append result
	if target == 0 {
		now := make([]int, len(*cur))
		copy(now, *cur)
		*result = append(*result, now)
		return
	}

	// end
	if target < 0 || idx >= len(cans) {
		return
	}

	curIdx := len(*cur)
	for target >= 0 {
		dfs(cans, idx+1, target, cur, result)
		*cur = append(*cur, cans[idx])
		target = target - cans[idx]
	}

	// clean
	*cur = (*cur)[:curIdx]
}

// @lc code=end

