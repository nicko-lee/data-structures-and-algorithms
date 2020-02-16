package main

import "fmt"

func IsPalindrome(str string) bool {
	// loop thru str from both ends at the same time comparing the 2
	for i, char := range str {
		if byte(char) != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	input := "abcdcba"
	fmt.Println(IsPalindrome(input))
}
