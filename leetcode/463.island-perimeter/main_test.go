package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIslandPerimeter(t *testing.T) {
	type test struct {
		name         string
		input2DArray [][]int
		expected     int
	}

	tests := []test{
		{name: "test one", input2DArray: [][]int{
			{0, 1, 0, 0},
			{1, 1, 1, 0},
			{0, 1, 0, 0},
			{1, 1, 0, 0},
		}, expected: 16},
		{name: "test two", input2DArray: [][]int{
			{0, 1, 0, 0},
			{1, 1, 1, 0},
			{0, 0, 1, 0},
			{1, 1, 1, 0},
		}, expected: 18},
		{name: "full land island", input2DArray: [][]int{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		}, expected: 16},
		{name: "no land island", input2DArray: [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}, expected: 0},
	}

	for _, testCase := range tests {
		actual := islandPerimeter(testCase.input2DArray)
		fmt.Println(actual)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Fatalf("test case: %s failed. expected: %v, got: %v", testCase.name, testCase.expected, actual)
		}
	}
}

func TestMain(t *testing.T) {
	main()
}
