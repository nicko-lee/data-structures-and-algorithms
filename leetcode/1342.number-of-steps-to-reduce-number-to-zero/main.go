package main

import "fmt"

func numberOfSteps(num int) int {
	steps := 0
	// reduce a number to 0. But can only divide 2 or subtract 1
	for num > 0 {
		// if even number
		if num%2 == 0 {
			num = num / 2
		} else {
			num--
		}
		steps++
	}
	return steps
}

func main() {
	input := 14
	fmt.Println(numberOfSteps(input))
}
