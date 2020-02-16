package main

import "fmt"

func countCharacters(words []string, chars string) int {

	// range through all possible words
	for _, v := range words {
		fmt.Println(v)
	}
	// checking for each if it can be formed from chars - create helper func dedicated to this
	// if yes keep track of the length and accumulate it in a result var

	return 9
}

// create helper func
func makeableWord(word string, chars string) bool {
	
	return nil
}
}

func main() {
	words := []string{"cat", "bt", "hat", "tree"}
	chars := "atach"

	out := countCharacters(words, chars)
	fmt.Println(out)
}
