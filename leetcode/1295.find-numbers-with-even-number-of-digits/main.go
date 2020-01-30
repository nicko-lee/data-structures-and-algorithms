package main

import (
	"fmt"
	"strconv"
)

func findNumbers(nums []int) int {
	var count int

	for _, v := range nums {
		// this means it is even number: https://www.dotnetperls.com/odd-go
		if len(strconv.Itoa(v))%2 == 0 {
			count++
		}
	}

	return count
}

func main() {
	inputNums := []int{555, 901, 482, 1771} // func input

	fmt.Println(findNumbers(inputNums))
}
