package main

import "fmt"

func BubbleSort(array []int) []int {
	sorted := false
	for !sorted {
		// do one pass
		for i := 0; i <= len(array)-2; i++ {
			current := array[i]
			next := array[i+1]
			if current > next {
				array[i] = next
				array[i+1] = current
			} else {
				// check if anything changed in this pass. If nothing changed. Means it is sorted
				sorted = true
				break
			}
		}
	}

	return array
}

func main() {
	input := []int{8, 5, 2, 9, 5, 6, 3}
	fmt.Println(BubbleSort(input))
}
