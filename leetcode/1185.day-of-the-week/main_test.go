package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDayOfTheYear(t *testing.T) {
	type test struct {
		name     string
		day      int
		month    int
		year     int
		expected string
	}

	tests := []test{
		{name: "test one", day: 31, month: 8, year: 2019, expected: "Saturday"},
		{name: "test two", day: 18, month: 7, year: 1999, expected: "Sunday"},
		{name: "test three", day: 15, month: 8, year: 1993, expected: "Sunday"},
		{name: "test three", day: 10, month: 10, year: 1990, expected: "Wednesday"},
	}

	for _, testCase := range tests {
		actual := dayOfTheWeek(testCase.day, testCase.month, testCase.year)
		fmt.Println(actual)
		if actual != testCase.expected {
			t.Fatalf("test case: %s failed. expected: %v, got: %v", testCase.name, testCase.expected, actual)
		}
	}
}

func TestIsLeapYear(t *testing.T) {
	type test struct {
		name      string
		inputYear int
		expected  bool
	}

	tests := []test{
		{name: "test %400 case", inputYear: 2000, expected: true},
		{name: "test two", inputYear: 1980, expected: true},
		{name: "test three", inputYear: 1972, expected: true},
		{name: "test odd number #1", inputYear: 2001, expected: false},
		{name: "test odd number #2", inputYear: 1977, expected: false},
	}

	for _, testCase := range tests {
		actual := isLeapYear(testCase.inputYear)
		fmt.Println(actual)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Fatalf("test case: %s failed. expected: %v, got: %v", testCase.name, testCase.expected, actual)
		}
	}
}

// testing it against this list of 32 known leap years between 1971 - 2100
// https://miniwebtool.com/leap-years-list/?start_year=1971&end_year=2100
func TestIsLeapYear2(t *testing.T) {
	// setup data structs for this test
	type resultStruct struct {
		listOfLeapYears []int
		total           int
	}

	type test struct {
		name     string
		input    []int
		expected resultStruct
	}

	// instantiate a data struct with test data
	testz := test{
		name: "test case",
		input: []int{1972, 1976, 1980, 1984, 1988, 1992, 1996, 2000, 2004, 2008, 2012,
			2016, 2020, 2024, 2028, 2032, 2036, 2040, 2044, 2048, 2052, 2056, 2060, 2064, 2068, 2072, 2076,
			2080, 2084, 2088, 2092, 2096},
		expected: resultStruct{
			listOfLeapYears: []int{1972, 1976, 1980, 1984, 1988, 1992, 1996, 2000, 2004, 2008, 2012,
				2016, 2020, 2024, 2028, 2032, 2036, 2040, 2044, 2048, 2052, 2056, 2060, 2064, 2068, 2072, 2076,
				2080, 2084, 2088, 2092, 2096},
			total: 32,
		},
	}

	count := 0
	for _, eachYear := range testz.input {
		actualLeapYears := isLeapYear(eachYear)
		if actualLeapYears {
			count++
		}
		fmt.Println(actualLeapYears)
	}
	fmt.Println(count)
	if count != testz.expected.total {
		t.Fatalf("test case: %s failed. expected: %v, got: %v", testz.name, testz.expected.total, count)
	}
}

func TestMain(t *testing.T) {
	main()
}
