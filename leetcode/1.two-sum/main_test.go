package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type test struct {
		name        string
		inputArray  []int
		inputTarget int
		expected    []int
	}

	tests := []test{
		{name: "test one", inputArray: []int{2, 7, 11, 15}, inputTarget: 9, expected: []int{0, 1}},
		{name: "test two", inputArray: []int{2, 7, 10, 15}, inputTarget: 12, expected: []int{0, 2}},
		{name: "first index is not 0", inputArray: []int{3, 2, 4}, inputTarget: 6, expected: []int{1, 2}},
		{name: "test four", inputArray: []int{3, 2, 4, 10, 9}, inputTarget: 19, expected: []int{3, 4}},
	}

	for _, testCase := range tests {
		actual := twoSum(testCase.inputArray, testCase.inputTarget)
		fmt.Println(actual)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Fatalf("test case: %s failed. expected: %v, got: %v", testCase.name, testCase.expected, actual)
		}
	}
}

func TestTwoSumDifferentImpl(t *testing.T) {
	type test struct {
		name        string
		inputArray  []int
		inputTarget int
		expected    []int
	}

	tests := []test{
		{name: "test one", inputArray: []int{2, 7, 11, 15}, inputTarget: 9, expected: []int{0, 1}},
		{name: "test two", inputArray: []int{2, 7, 10, 15}, inputTarget: 12, expected: []int{0, 2}},
		{name: "first index is not 0", inputArray: []int{3, 2, 4}, inputTarget: 6, expected: []int{1, 2}},
		{name: "test four", inputArray: []int{3, 2, 4, 10, 9}, inputTarget: 19, expected: []int{3, 4}},
	}

	for _, testCase := range tests {
		actual := twoSumDifferentImpl(testCase.inputArray, testCase.inputTarget)
		fmt.Println(actual)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Fatalf("test case: %s failed. expected: %v, got: %v", testCase.name, testCase.expected, actual)
		}
	}
}
