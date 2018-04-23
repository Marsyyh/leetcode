package main

import "fmt"

// Get max range for each iteration, if end = max range means next step
func main() {
	test := makeRange(1, 25000)
	res := jump(test)
	fmt.Print(res)
}

func jump(nums []int) int {
	s, e, max := 0, 0, 0
	if len(nums) < 2 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		max = twoValueMax(max, i+nums[i])
		if max >= len(nums)-1 {
			break
		}
		if i == e {
			s++
			e = max
		}
	}
	return s + 1
}

func twoValueMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// for 25000 length, stop
// Time O(n2), space O(n)
func jump2(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	step := 1
	position := map[int]bool{
		0: true,
	}
	for step <= len(nums) {
		temSlice := make(map[int]bool)
		for i := range position {
			result := makeRange(i+1, i+1+nums[i])
			for _, j := range result {
				if _, ok := temSlice[j]; !ok {
					temSlice[j] = true
				}
			}
		}
		if _, ok := temSlice[len(nums)-1]; ok {
			break
		}
		position = temSlice
		step++
	}
	return step
}

func makeRange(min, max int) []int {
	a := make([]int, max-min)
	for i := range a {
		a[i] = min + i
	}
	return a
}
