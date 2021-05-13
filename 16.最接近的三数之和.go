/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
/*
	all 3: O(n^3)

	twoSumClosest:
	sorted:
	-3 -2 -1  0 1 2 3
	l    			r
	if r + l > target: r-- update sum
	if r + l < target: l++ update sum
*/
func threeSumClosest(nums []int, target int) int {
	// len(sums) gurante >= 3
	sort.Ints(nums)
	sum := -9999
	for i := 0; i < len(nums)-1; i++ {
		twoSum := twoSumClosest(nums, i+1, len(nums)-1, target-nums[i])
		if abs(sum-target) > abs(nums[i]+twoSum-target) {
			sum = nums[i] + twoSum
		}
	}
	return sum
}

func twoSumClosest(nums []int, l, r, target int) int {
	sum := -9999
	for l < r {
		cur := nums[l] + nums[r]
		if cur < target {
			l++
		} else {
			r--
		}
		if abs(sum-target) > abs(cur-target) {
			sum = cur
		}
	}
	return sum
}

func abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

// @lc code=end

