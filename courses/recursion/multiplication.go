package main

func mult_iter(a int, b int) int {
	result := 0
	for b > 0 { // won't enter loop if b = 0 you will simply exit the function and return 0
		result += a // sum up a for b number of times
		b--         // reduce iteration var to prevent infinite loop
	}
	return result
}

func mult_recur(a int, b int) int {
	// specify base case
	if b == 1 {
		return a
	} else {
		return a + mult_recur(a, b-1)
	}
}
