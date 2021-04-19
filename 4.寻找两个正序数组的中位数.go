/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
/*
	1, 2, 5
	3, 4, 6


	2, 5 -> 3.5
	3, 4 -> 3.5

*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := []int{}
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			n = append(n, nums1[i])
			i++
		} else {
			n = append(n, nums2[j])
			j++
		}
	}

	if i < len(nums1) {
		n = append(n, nums1[i:]...)
	}

	if j < len(nums2) {
		n = append(n, nums2[j:]...)
	}

	return findMedian(n)
}

func findMedian(nums []int) float64 {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l%2 == 0 {
		return (float64(nums[l/2-1]) + float64(nums[l/2])) / 2
	}
	return float64(nums[l/2])
}

// @lc code=end

