package main

import "fmt"

func islandPerimeter(grid [][]int) int {
	perimeterCount := 0

	// access individual rows
	for i, row := range grid {
		// access individual elements
		for i2, element := range row {
			// for each item check if it is land or water. If it is water do nothing
			if element == 1 {
				// for the following checks only add 1 to count if it is water or out of bounds
				// check left
				if i2-1 < 0 { // out of bounds check (prevents an out of bounds error)
					perimeterCount++
				} else if row[i2-1] == 0 {
					perimeterCount++
				}

				// check right (note len(row) is out of bounds)
				// see https://play.golang.org/p/uz2xI0pWHGm
				if i2+1 == len(row) { // out of bounds check
					perimeterCount++
				} else if row[i2+1] == 0 {
					perimeterCount++
				}

				// check up (note i is index for row, i2 is index for element in row)
				if i-1 < 0 || grid[i-1][i2] == 0 {
					perimeterCount++
				}

				// check down
				if i+1 == len(grid) || grid[i+1][i2] == 0 {
					perimeterCount++
				}
			}
		}
	}
	return perimeterCount
}

func main() {
	// can access by [row][column] notation
	input := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	fmt.Println(islandPerimeter(input))
}

/* -------- NOTES -----------

This is how you traverse a 2D slice/array in Go:

	// Loop over slices in animals. Access rows
	// Source: https://www.dotnetperls.com/2d-go
    for i := range animals {
        fmt.Printf("Row: %v\n", i)
        fmt.Println(animals[i])
    }

The above gives you just each row. But then you need to loop again to access each individual element in the row

	// access every individual item in order from top left to bottom right (just like reading through a page)
	itemCount := 0
	for _, row := range input {
		for _, element := range row {
			itemCount++
			fmt.Printf("Element number: %v\n", itemCount)
			fmt.Println(element)
		}
	}

*/
