package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDefangIP(t *testing.T) {
	type test struct {
		name      string
		inputAddr string
		expected  string
	}

	tests := []test{
		{name: "test one", inputAddr: "192.168.1.1", expected: "192[.]168[.]1[.]1"},
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
