package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetAndStoreData(t *testing.T) {
	type test struct {
		name      string
		inputList []int
		expectedType LinkedList
		expectedReturn  []int
	}

	tests := []test{
		{name: "test one", inputList: []int{1,0,1}, expected: LinkedList{Head: }},
		{name: "test one", inputAddr: "255.100.50.0", expected: "255[.]100[.]50[.]0"},
	}

	for _, testCase := range tests {
		actual := defangIPaddr(testCase.inputAddr)
		fmt.Println(actual)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Fatalf("test case: %s failed. expected: %v, got: %v", testCase.name, testCase.expected, actual)
		}
	}
}
