/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	size := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[size] = nums[i]
			size++
		}
	}
	return size
}

// @lc code=end

