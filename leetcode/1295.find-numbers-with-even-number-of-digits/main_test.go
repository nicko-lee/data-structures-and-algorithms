package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindNums(t *testing.T) {
	type test struct {
		name       string
		inputArray []int
		expected   int
	}

	tests := []test{
		{name: "test one", inputArray: []int{2, 7, 11, 15}, expected: 2},
		{name: "test two", inputArray: []int{12, 345, 2, 6, 7896}, expected: 2},
		{name: "test two", inputArray: []int{12, 345, 2, 6, 7896, 777788}, expected: 3},
	}

	for _, testCase := range tests {
		actual := findNumbers(testCase.inputArray)
		fmt.Println(actual)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Fatalf("test case: %s failed. expected: %v, got: %v", testCase.name, testCase.expected, actual)
		}
	}
}
