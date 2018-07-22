package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	origNums, nums []int
}

func Constructor(nums []int) Solution {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	return Solution{tmp, nums}
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

func main() {
	nums := []int{1, 2, 3, 4, 5}
	obj := Constructor(nums)
	param_2 := obj.Shuffle()
	fmt.Println(param_2)
	param_1 := obj.Reset()
	fmt.Println(param_1)
}
