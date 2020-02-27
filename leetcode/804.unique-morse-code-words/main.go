package main

import "fmt"

func uniqueMorseRepresentations(words []string) int {
	set := make(map[string]bool) // new empty set to enable check for membership
	for _, word := range words {
		set[convertToMorse(word)] = true // chuck Morse Represenation into set
	}
	size := len(set) // size of set (represents unique morse representations)
	return size
}

// helper function for above func
func convertToMorse(word string) string {
	m := createMorseMap()
	morseCode := ""

	for _, v := range word {
		val, _ := m[string(v)]
		morseCode += val
	}

	return morseCode
}

func createMorseMap() map[string]string {
	morseCode := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.",
		"---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	m := make(map[string]string)

	var alphabet rune
	alphabet = 97 // initialise for "a"
	for _, v := range morseCode {
		m[string(alphabet)] = v
		alphabet += 1
	}

	return m
}

func main() {
	m := make(map[string]string)
	m["a"] = "sakai"

	fmt.Println(m)

	// fmt.Println(createMorseMap())

	inputWords := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(inputWords))

}

/*
-----Notes-----

High level steps:
	1. Convert array to map that maps alphabets a-z to corresponding Morse Code -> O(1)
	2. Range over input words -> O(N)
		3. For each word range over each character and look up in map -> O(N)
			4. Build Morse String
	5. Check for unique morse representations

The zero value of a map is nil. A nil map has no keys.
Moreover, any attempt to add keys to a nil map will result in a runtime error.
https://www.callicoder.com/golang-maps/

So when I did the following:
	func main() {
		var m map[string]string
		m["a"] = "sakai"

		fmt.Println(m)
	}

And ran it -> it panicked out saying "assignment to entry in nil map"

It is, therefore, necessary to initialize a map before adding items to it.

2 ways to initialize a map in Go:
	1.Using built in make() func
	2. Using a map literal

For convenience here is the mapping:
	map[a:.- b:-... c:-.-. d:-.. e:. f:..-. g:--. h:.... i:.. j:.--- k:-.- l:.-.. m:-- n:-.
	o:--- p:.--. q:--.- r:.-. s:... t:- u:..- v:...- w:.-- x:-..- y:-.-- z:--..]

Very simple Set implementation! This was a really useful data structure for this question
https://yourbasic.org/golang/implement-set/

Nice example on easily converting byte (uint8) and rune (int32) from a number into a string
https://play.golang.org/p/1aXDgrckHKs
*/
