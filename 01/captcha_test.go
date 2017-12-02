package main

import (
	"testing"
)

func TestSum(t *testing.T){
	type testCase struct {
		in []int
		expected int
	}

	testCases := []testCase {
		{in: []int{1, 1, 2, 2}, expected: 3},
		{in: []int{1, 1, 1, 1}, expected: 4},
		{in: []int{1, 2, 3, 4}, expected: 0},
		{in: []int{9,1,2,1,2,1,2,9}, expected: 9},
	}

	for i, tc := range testCases{
		if sum(tc.in) != tc.expected{
			t.Errorf("Test %d failed. Expected %d, got %d", i, tc.expected, sum(tc.in))
		}
	}

}
func TestOtherSum(t *testing.T){
	type testCase struct {
		in []int
		expected int
	}

	testCases := []testCase {
		{in: []int{1, 1, 2, 2}, expected: 6},
		{in: []int{1, 2, 1, 2}, expected: 0},
		{in: []int{1, 2, 3, 4, 2, 5}, expected: 4},
		{in: []int{1,2,3,1,2,3}, expected: 12},
		{in: []int{1,2,1,3,1,4,1,5}, expected: 4},
	}

	for i, tc := range testCases{
		if sum(tc.in) != tc.expected{
			t.Errorf("Test %d failed. Expected %d, got %d", i, tc.expected, sum(tc.in))
		}
	}
}
