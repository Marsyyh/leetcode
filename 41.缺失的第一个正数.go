/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
/*

	first missing must in [1, N+1]
	because if 1 to N exist, i will be N+1

	1. first time move all not in [1, N] to 0
	2. orderly check which hole is first 0
*/
func firstMissingPositive(nums []int) int {
	// first time
	l := len(nums)
	for i := 0; i < l; i++ {
		for nums[i] != 0 && nums[i] != i+1 {
			cur := nums[i]
			// end since find beyond boundary
			if cur > l || cur < 1 || nums[cur-1] == nums[i] {
				nums[i] = 0
				break
			}
			// swap (cur, nums[cur-1]), go find next
			nums[i], nums[cur-1] = nums[cur-1], nums[i]
		}
	}

	// second time
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			return i + 1
		}
	}
	return l + 1
}

// @lc code=end

