/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	size := 0
	for i := 0; i < len(nums); i++ {
		// shorten same elements
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
		nums[size] = nums[i]
		size++
	}
	return size
}

// @lc code=end

