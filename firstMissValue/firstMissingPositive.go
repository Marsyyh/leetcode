package main

import "fmt"

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

func main() {
	v := []int{3, 4, -1, 1}
	r := firstMissingPositive2(v)
	fmt.Println(r)
}

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
