/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
/*
	find two values then two sum
*/
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	re := [][]int{}
	for i := 0; i < len(nums)-1; i++ {
		// same value
		if i < len(nums)-3 && nums[i+1] == nums[i] {
			two := twoSum(nums, i+2, len(nums)-1, target-2*nums[i])
			for _, twoR := range two {
				cur := append([]int{nums[i], nums[i]}, twoR...)
				re = append(re, cur)
			}
			for i < len(nums)-1 && nums[i] == nums[i+1] {
				i++
			}
		}

		for j := i + 1; j < len(nums)-1; j++ {
			// same value
			if j < len(nums)-2 && nums[j+1] == nums[j] {
				res := target - nums[i] - 2*nums[j]
				for k := j + 2; k < len(nums); k++ {
					if nums[k] == res {
						re = append(re, []int{nums[i], nums[j], nums[j], res})
						break
					}
				}
				for j < len(nums)-1 && nums[j] == nums[j+1] {
					j++
				}
			}

			two := twoSum(nums, j+1, len(nums)-1, target-nums[i]-nums[j])
			for _, twoR := range two {
				cur := append([]int{nums[i], nums[j]}, twoR...)
				re = append(re, cur)
			}
		}
	}
	return re
}

// sorted two sum, left & right pointer
func twoSum(nums []int, l, r, target int) [][]int {
	re := [][]int{}
	for l < r {
		// left, right same
		if nums[l] == nums[r] {
			if nums[l]+nums[r] == target {
				re = append(re, []int{nums[l], nums[r]})
			}
			return re
		}

		// find target value, append
		if nums[l]+nums[r] == target {
			re = append(re, []int{nums[l], nums[r]})
		}

		// left move to next non same
		if nums[l]+nums[r] < target {
			// handle same l value eq target
			if l+1 < len(nums) && nums[l] == nums[l+1] && 2*nums[l] == target {
				re = append(re, []int{nums[l], nums[l]})
				return re
			}
			// moving l to next non same
			for l < len(nums)-1 && nums[l+1] == nums[l] {
				l++
			}
			l++
		} else {
			// right move to next non same
			// handle same r value eq target
			if r-1 >= 0 && nums[r] == nums[r-1] && 2*nums[r] == target {
				re = append(re, []int{nums[r], nums[r]})
				return re
			}

			// moving r to next non same
			for r > 0 && nums[r-1] == nums[r] {
				r--
			}
			r--
		}
	}
	return re
}

// @lc code=end

