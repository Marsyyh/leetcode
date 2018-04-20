package main

import "fmt"

// loop, put num > 0 and < size into correct spot in new array, check first value/spot not match element
// Time O(n), space O(n)
func firstMissingPositive(nums []int) int {
	l := len(nums)
	h := make([]int, l)
	for _, n := range nums {
		if n > 0 && n <= l {
			h[n-1] = n
		}
	}
	for i, n := range h {
		if n-1 != i {
			return i + 1
		}
	}
	return l + 1
}

// instead put into new array, now do swap.
// 1. after swap, should re-check new value
// 2. check if current spot value = swap value, skip it, otherwise it will go infinite loop
// Time O(n), Space O(1)
func firstMissingPositive2(nums []int) int {
	l := len(nums)
	for i := 0; i < l; {
		if nums[i] > 0 && nums[i] <= l && nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
			continue
		}
		i++
	}

	for i, n := range nums {
		if n != i+1 {
			return i + 1
		}
	}
	return l + 1
}

func main() {
	v := []int{3, 4, -1, 1}
	r := firstMissingPositive2(v)
	fmt.Println(r)
}
