/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	re := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		n := nums[i]
		// same value
		if i+1 < len(nums) && n == nums[i+1] {
			// find three 0
			if n == 0 && nums[i+1] == 0 && i+2 < len(nums) && nums[i+2] == 0 {
				re = append(re, []int{0, 0, 0})
			}

			// Set i value && find two same
			for j := i + 1; j < len(nums); j++ {
				// set i at exact change
				if nums[j-1] == n && (nums[j] != n || j == len(nums)-1) {
					i = j - 1
				}

				if nums[j] != 0 && nums[j] == -2*n {
					re = append(re, []int{n, n, -2 * n})
					break
				}
			}
		}

		nCouple := twoSum(nums, -n, i)
		for _, c := range nCouple {
			re = append(re, append(c, n))
		}
	}
	return re
}

func twoSum(nums []int, sum, idx int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	re := [][]int{}
	m := make(map[int]int)
	for i := idx + 1; i < len(nums); i++ {
		n := nums[i]
		_, ok := m[n]
		// if not in, add to list
		if !ok {
			m[sum-n] = 0
			continue
		}
		if m[sum-n] == 0 {
			re = append(re, []int{n, sum - n})
			m[sum-n]++
		}
	}

	return re
}

// @lc code=end
