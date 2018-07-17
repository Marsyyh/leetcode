package main

import "fmt"

func main() {
	test := []int{1, 3, 5, 7}
	target := 6
	res := searchInsert(test, target)
	fmt.Print(res)
}

// Binary search
// left <= right : left = mid + 1; right = mid - 1. Target one element in the last nums[mid] == target
// left < right: left = mid + 1; right = mid. Target
// left < right - 1: left = mid, right = mid.
// Target will fall into [left, right], unless target < nums[0] pr target > nums[len(nums) - 1]
func searchInsert(nums []int, target int) int {
	var left, right int
	for left, right = 0, len(nums)-1; left < right-1; {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}
