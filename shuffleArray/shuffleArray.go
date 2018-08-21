package shuffleArray

import (
	"math/rand"
)

type Solution struct {
	origNums, nums []int
}

func Constructor(nums []int) Solution {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	return Solution{nums, tmp}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.origNums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	for i, _ := range this.nums {
		index := rand.Intn(len(this.nums) - i)
		this.nums[i], this.nums[index+i] = this.nums[index+i], this.nums[i]
	}
	return this.nums
}
