/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
/*
	alphabetical order
	abc < acb < bac < bca < cba

	next order:
	1. from right to left, find most right two place to swap
		left place must be right most
		right place must smallest
	2. reverse the right part to ascending order since it's smallest
*/
func nextPermutation(nums []int) {
	// find left and right
	left, right := -1, -1
	// iterate from right
	for i := len(nums) - 2; i >= 0 && left == -1; i-- {
		// find right smallest than nums[i], if not keep iterate
		for j := i + 1; j < len(nums); j++ {
			// find it
			if nums[i] < nums[j] {
				left = i
				if right == -1 {
					right = j
				} else if nums[j] < nums[right] {
					right = j
				}
			}
		}
	}

	// swap
	if left != -1 {
		nums[left], nums[right] = nums[right], nums[left]
	}

	// reverse assending order
	reverse(nums, left)
}

func reverse(nums []int, idx int) {
	for i := idx + 1; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

// @lc code=end

