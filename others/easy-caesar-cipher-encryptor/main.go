package main

import "fmt"

func CaesarCipherEncryptor(str string, key int) string {
	encrypted := ""

	// used http://www.asciitable.com/
	for _, v := range str {
		shiftedChar := v + rune(key)
		// also need to take into account edge cases of very big keys
		var multipleOf26 int
		if key > 26 {
			multipleOf26 = key / 26
		} else {
			multipleOf26 = 1
		}

		// take into account wrapping past z
		if shiftedChar > 122 {
			shiftedChar -= (26 * rune(multipleOf26))
		}
		encrypted += string(shiftedChar) // shift each char in string by 2
	}
	return encrypted
}

func main() {
	input := "apple"
	key := 2
	fmt.Println(rune('a')) // note in Go "a" is different to 'a' (syntactic sugar for Unicode value)
	fmt.Println(CaesarCipherEncryptor(input, key))
}
