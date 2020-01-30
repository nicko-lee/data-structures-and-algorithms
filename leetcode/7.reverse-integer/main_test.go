package main

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type test struct {
		name     string
		inputInt int
		expected int
	}

	tests := []test{
		{name: "test one", inputInt: 123, expected: 321},
		{name: "test signed int", inputInt: -123, expected: -321},
		{name: "test ending with zero", inputInt: 120, expected: 21},
		{name: "test four", inputInt: 1231, expected: 1321},
	}

	for _, testCase := range tests {
		actual := reverse(testCase.inputInt)
		fmt.Println(actual)
		if testCase.expected != actual {
			t.Fatalf("test case: %s failed. expected: %v, got: %v", testCase.name, testCase.expected, actual)
		}
	}
}
