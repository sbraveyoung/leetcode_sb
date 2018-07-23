package Shuffle

import (
	"math/rand"
)

type Solution struct {
	origin []int
}

func Constructor(nums []int) Solution {
	return Solution{
		origin: nums,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.origin
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	length := len(this.origin)
	m := map[int]int{}
	ret := []int{}
	for {
		if len(ret) == length {
			break
		}
		r := rand.Intn(length)
		if _, ok := m[r]; ok {
			continue
		}
		m[r] = r
		ret = append(ret, this.origin[r])
	}
	return ret
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
