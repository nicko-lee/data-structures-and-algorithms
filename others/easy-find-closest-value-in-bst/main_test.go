package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)



var test = NewBST(100, 5, 15, 5, 2, 1, 22, 1, 1, 3, 1, 1, 502, 55000, 204, 205, 207, 206, 208, 203, -51, -403, 1001, 57, 60, 4500)

// case exact match at head
func TestCase1(t *testing.T) {
	output := test.FindClosestValue(100)
	expected := 100
	require.Equal(t, expected, output)
}

// case exact match not at head node
func TestCase2(t *testing.T) {
	output := test.FindClosestValue(208)
	expected := 208
	require.Equal(t, expected, output)
}

// case exact match at leaf node
func TestCase3(t *testing.T) {
	output := test.FindClosestValue(4500)
	expected := 4500
	require.Equal(t, expected, output)
}

func TestCase4(t *testing.T) {
	output := test.FindClosestValue(4501)
	expected := 4500
	require.Equal(t, expected, output)
}

func TestCase5(t *testing.T) {
	output := test.FindClosestValue(-70)
	expected := -51
	require.Equal(t, expected, output)
}

func TestCase6(t *testing.T) {
	output := test.FindClosestValue(2000)
	expected := 1001
	require.Equal(t, expected, output)
}

func TestCase7(t *testing.T) {
	output := test.FindClosestValue(6)
	expected := 5
	require.Equal(t, expected, output)
}

func TestCase8(t *testing.T) {
	output := test.FindClosestValue(30000)
	expected := 55000
	require.Equal(t, expected, output)
}

func TestCase9(t *testing.T) {
	output := test.FindClosestValue(-1)
	expected := 1
	require.Equal(t, expected, output)
}

func TestCase10(t *testing.T) {
	output := test.FindClosestValue(29751)
	expected := 55000
	require.Equal(t, expected, output)
}

func TestCase11(t *testing.T) {
	output := test.FindClosestValue(29749)
	expected := 4500
	require.Equal(t, expected, output)
}

// func TestMain(t *testing.T) {
// 	main()
// 	stringify(test, 100)
// }
