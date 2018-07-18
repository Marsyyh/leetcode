package main

import "fmt"

func main() {
	test := []int{1}
	target := 0
	res := searchRange(test, target)
	fmt.Print(res)
}

// binary search left and right
func searchRange(nums []int, target int) []int {
	var left, right int
	result := []int{-1, -1}
	if nums == nil || len(nums) == 0 {
		return result
	}
	// find most left ==
	for left, right = 0, len(nums)-1; left < right; {
		mid := left + (right-left)/2
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if nums[left] == target {
		result[0] = left
	}

	// find most right ==
	// right need to end with len(nums)
	// because if target in most right, left eventually will == right, so we cover this case
	for left, right = 0, len(nums); left < right; {
		mid := left + (right-left)/2
		if target < nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left > 0 && nums[left-1] == target {
		result[1] = left - 1
	}
	return result
}

// find the element then goes to most left, right
// however, worst case Time is O(n) because same elements
func searchRange1(nums []int, target int) []int {
	var left, right int
	result := []int{-1, -1}
	if nums == nil || len(nums) == 0 {
		return result
	}
	index := binarySearch(nums, target)
	if index == -1 {
		return result
	}
	// find the most left, short cut if left out of bundary
	for left = index; left >= 0 && nums[left] == target; {
		left--
	}

	// find the most right, short cut if right out of bundary
	for right = index; right < len(nums) && nums[right] == target; {
		right++
	}

	// beacuse we minus and add after last == or bundary, we need revert them back
	result[0] = left + 1
	result[1] = right - 1
	return result
}

func binarySearch(nums []int, target int) int {
	var mid int
	for left, right := 0, len(nums)-1; left <= right; {
		mid = left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
