package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected bool
	}

	tests := []test{
		{name: "test one", input: "abcdcba", expected: true},
		{name: "test two", input: "anna", expected: true},
		{name: "test three", input: "civic", expected: true},
		{name: "test four", input: "reddit", expected: false},
		{name: "test five", input: "kfc", expected: false},
		{name: "test six", input: "annabella", expected: false},
	}

	for _, testCase := range tests {
		actual := IsPalindrome(testCase.input)
		fmt.Println(actual)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Fatalf("test case: %s failed. expected: %v, got: %v", testCase.name, testCase.expected, actual)
		}
	}
}
