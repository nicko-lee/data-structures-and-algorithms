package main

import (
	"fmt"
	"testing"
)

func TestMakeableWord(t *testing.T) {
	type test struct {
		name     string
		word     string
		chars    string
		expected bool
	}

	tests := []test{
		{name: "test one", word: "cat", chars: "atach", expected: true},
		{name: "test two", word: "bt", chars: "atach", expected: false},
		{name: "test three", word: "tree", chars: "atach", expected: false},
		{name: "test four", word: "hat", chars: "atach", expected: true},
	}

	for _, testCase := range tests {
		actual := makeableWord(testCase.word)
		fmt.Println(actual)
		if testCase.expected != actual {
			t.Fatalf("test case: %s failed. expected: %v, got: %v", testCase.name, testCase.expected, actual)
		}
	}
}
