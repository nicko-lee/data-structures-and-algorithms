package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaximum69Number(t *testing.T) {
	type test struct {
		name     string
		input    int
		expected int
	}

	tests := []test{
		{name: "test one", input: 9669, expected: 9969},
		{name: "test two", input: 9996, expected: 9999},
		{name: "test three", input: 9999, expected: 9999},
	}

	for _, testCase := range tests {
		actual := maximum69Number(testCase.input)
		fmt.Println(actual)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Fatalf("test case: %s failed. expected: %v, got: %v", testCase.name, testCase.expected, actual)
		}
	}
}
