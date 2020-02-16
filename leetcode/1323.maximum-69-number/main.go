package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
super naive brute force solution
pseudocode:
1. convert the input data to a form where I can access individual digits
	1.1 convert int to string
	1.2 now I can range thru string
	1.3 pluck out each individual digit with helper digit() func and push to a stack (just a slice of int)
2. process each element in stack checking for the first 6 and building the return string using concatenation digit by digit
	1.1 then converting back to an int

Highly inefficient but works. So many conversions back and forth. Uses two loops.
Good example to show how you can do the same thing with a one-liner
But at least I did learn how to work with stacks
*/
func maximum69Number(num int) int {
	// simply look for the first non-9 digit and change it to 9
	stringNum := strconv.Itoa(num)
	var stack []int

	// range through the number each digit by digit
	for i, _ := range stringNum {
		// plucking out each individual digit starting from the far right
		d := digit(num, i+1)
		// adding to a stack (cos it is FIFO structure)
		// source: https://yourbasic.org/golang/implement-stack/
		stack = append(stack, d) // Push
	}
	// ---- by this point all we did is "reshape" the input data so we can pluck out each individual digit ----

	// setting up some result vars for the next loop
	s := ""
	result := ""
	var swapped bool
	// go thru each element in stack
	for len(stack) > 0 {
		n := len(stack) - 1 // Index of top element
		// check for the first 6 and swap for a 9
		if stack[n] == 6 && !swapped {
			stack[n] = 9
			swapped = true
		}
		s = strconv.Itoa(stack[n])
		stack = stack[:n] // Pop (i.e. the array is now everything but for the last element stack[n])
		result = result + s
	}

	// convert back to int and return
	resultInt, _ := strconv.Atoi(result)
	return resultInt
}

// source: https://stackoverflow.com/questions/46753736/extract-digits-at-a-certain-position-in-a-number
// starts counting from the far left. Start counting from 1 not index 0
// currently if you use this to print out say 9969 it will print as 9699
// try and change it to print from the start
func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}

// smart alternative solution: https://leetcode.com/problems/maximum-69-number/discuss/484405/Go-Need-to-replace-only-once
// strings.Replace solves exactly that it takes the first non 9 number and replaces that
func maximum69NumberV2(num int) int {
	str := strconv.Itoa(num) // converts the num to a string
	str = strings.Replace(str, "6", "9", 1)
	result, _ := strconv.Atoi(str)
	return result
}

func main() {
	// fmt.Println(maximum69NumberV2(9969))
	fmt.Println(maximum69Number(9969))
}

/* ---------------- NOTES ------------------

very well explained here:
//https://leetcode.com/problems/maximum-69-number/discuss/486969/GoGolang-Simple-2-line-solution-with-strings.Replace-(100-Runtime100-Memory)

*/
