package main

import (
	"fmt"
)

func defangIPaddr(address string) string {
	// source: https://stackoverflow.com/questions/18130859/how-can-i-iterate-over-a-string-by-runes-in-go
	var defanged string

	// note that when u range the string - each individual char is not a string but a rune interestingly enough
	for _, char := range address {
		if char == rune('.') {
			defanged += string("[.]")
		} else {
			defanged += string(char)
		}
	}
	return defanged
}

func main() {
	fmt.Println("sakai") // prove program works
	address := "1.1.1.1" // input IP address

	fmt.Println(defangIPaddr(address)) // output
}

// other solutions
// wow this guy did it in one line lol - https://leetcode.com/problems/defanging-an-ip-address/discuss/464480/Golang-solution
