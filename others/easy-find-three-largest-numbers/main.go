package main

import "fmt"

func FindThreeLargestNumbers(array []int) []int {
	topThree := []int{}

	for _, v1 := range array {
		// grab each item and compare with each item currently in topThree
		if len(topThree) < 3 {
			topThree = append(topThree, v1)
		} else {
			for i2, v2 := range topThree {
				if v1 > v2 {
					// the larger number takes the place of the smaller
					topThree[i2] = v1
				}
			}
		}
	}

	// sort return array

	return topThree
}

func main() {
	input := []int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7}
	fmt.Println(FindThreeLargestNumbers(input))
}
