/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
/*
	binary search

	find k
	if half > left: k in (half, right)
				  : k in (left, half)
*/
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	// find k
	l, r := 0, len(nums)-1
	for l < r {
		h := l + (r-l)/2
		if nums[h] > nums[l] {
			l = h
		} else {
			r = h
		}
	}

	// init search area
	// if no k
	if l == len(nums)-1 || nums[l] < nums[l+1] {
		l, r = 0, len(nums)-1
	} else {
		// if k
		if target >= nums[0] {
			l = 0
		} else {
			l, r = l+1, len(nums)-1
		}
	}

	for l <= r {
		h := l + (r-l)/2
		if target == nums[h] {
			return h
		}

		if target < nums[h] {
			r = h - 1
		} else {
			l = h + 1
		}
	}

	return -1
}

// @lc code=end

