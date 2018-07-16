package main

import "fmt"

func main() {
	test := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(test, target)
	fmt.Print(res)
}

// Using Map, we can do it in O(n) time
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		// check if pair exsit before add to map
		// to avoid check i != j (same value can't be use twice)
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return []int{}
}
