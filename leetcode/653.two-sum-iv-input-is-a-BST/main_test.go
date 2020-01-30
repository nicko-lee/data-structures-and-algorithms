package main

import (
	"fmt"
	"testing"
)

func TestSearchTree(t *testing.T) {
	type test struct {
		name      string
		BST       *TreeNode
		searchVal int
		expected  bool
	}

	tests := []test{
		{name: "test one", BST: InitTree(5, 3, 6, 2, 4, 7), searchVal: 3, expected: true},
		{name: "test two", BST: InitTree(5, 3, 6, 2, 4, 7), searchVal: 7, expected: true},
		{name: "test three", BST: InitTree(5, 3, 6, 2, 4, 7), searchVal: 14, expected: false},
		{name: "test four", BST: InitTree(5, 3, 6, 2, 4, 7), searchVal: 9, expected: false},
	}

	for _, testCase := range tests {
		actual := searchTree(testCase.searchVal, testCase.BST)
		fmt.Println(actual)
		if actual != testCase.expected {
			t.Fatalf("test case: %s failed. expected: %v, got: %v", testCase.name, testCase.expected, actual)
		}
	}
}
