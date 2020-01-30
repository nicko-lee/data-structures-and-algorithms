package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// loop each element of array systematically from left to right
	for i, v := range nums {
		sum := 0
		// inner loop that for each given element wiil sum with elements to its right to check if equals target
		for i2 := i; i2 < len(nums)-1; i2++ {
			sum = v + nums[i2+1]
			if sum == target {
				return []int{i, i2 + 1}
			}
		}
	}
	// default return value
	return []int{0, 0}
}

// source: https://leetcode.com/problems/two-sum/discuss/237852/golang-solution-for-two-sum-faster-than-100-memory-less-than-100
// rationale on why it is only N not N-squared compared to my "naive" algorithm above: https://coderbyte.com/algorithm/two-sum-problem
func twoSumDifferentImpl(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {

		// note this works cos in the map it is an inversion of value: index
		if v, found := m[target-v]; found {
			return []int{v, i}
		}

		// here we are basically "remembering" things we have already come across so we can go back and lookup here
		// taking advantage of faster searching of a map rather than walking through an entire array each time
		// why do "hash tables" have highly efficient lookup? - https://www.quora.com/Why-is-a-lookup-in-a-hash-data-structure-so-fast-O-1
		m[v] = i
	}
	return nil
}

func main() {
	fmt.Println("I ran")
	t := 9
	n := []int{2, 7, 11, 15}

	fmt.Println(twoSumDifferentImpl(n, t))
}
