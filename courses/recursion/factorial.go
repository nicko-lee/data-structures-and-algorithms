package main

func factorial_recur(n int) int {
	// base case
	if n == 1 {
		return 1
	} else {
		// how to reduce problem? Rewrite in terms of something simpler to reach base case
		return n * factorial_recur(n-1)
	}
}

func factorial_iter(n int) int {
	product := n
	for i := n - 1; i > 0; i-- {
		product *= i
	}
	return product
}
