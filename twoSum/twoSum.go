package twoSum

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
