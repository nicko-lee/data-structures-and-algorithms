package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCaesarCipherEncryptor(t *testing.T) {
	type test struct {
		name     string
		inputStr string
		inputKey int
		expected string
	}

	tests := []test{
		{name: "test one", inputStr: "abc", inputKey: 2, expected: "cde"},
		{name: "test two", inputStr: "apple", inputKey: 1, expected: "bqqmf"},
		{name: "test three: wrapping past z", inputStr: "xyz", inputKey: 2, expected: "zab"},
		{name: "test four: big shift", inputStr: "abc", inputKey: 57, expected: "fgh"},
	}

	for _, testCase := range tests {
		actual := CaesarCipherEncryptor(testCase.inputStr, testCase.inputKey)
		fmt.Println(actual)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Fatalf("test case: %s failed. expected: %v, got: %v", testCase.name, testCase.expected, actual)
		}
	}
}
