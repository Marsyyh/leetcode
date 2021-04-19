/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

/*
map 解法时间复杂度O(n), 空间复杂度O(n)
左右指针解法时间复杂度，排序O(nlogn), 找O(n), 空间排序O(logn） 查找O(1)
*/
// @lc code=start
func twoSum(nums []int, target int) []int {
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	for i < j {
		if target-nums[i] == nums[j] {
			return []int{i, j}
		}

		if target-nums[i] > nums[j] {
			i++
		} else {
			j--
		}
	}
	return []int{}
}

func twoMapSum(nums []int, target int) []int {
	m := make(map[int]int)
	f, s := -1, -1
	for i, n := range nums {
		if mi, ok := m[n]; ok {
			f, s = mi, i
			break
		}
		m[target-n] = i
	}
	return []int{f, s}
}

// @lc code=end

