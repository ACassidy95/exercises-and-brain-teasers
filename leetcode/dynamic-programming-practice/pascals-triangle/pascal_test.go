package pascalstriangle_test

import (
	pascalstriangle "dynamic-programming-practice/pascals-triangle"
	"reflect"
	"testing"
)

func TestGeneratePascalsTriangle(t *testing.T) {
	testCases := map[string]struct {
		numRows  int
		expected [][]int
	}{
		"Single Row": {
			numRows:  1,
			expected: [][]int{{1}},
		},
		"Two Rows": {
			numRows: 2,
			expected: [][]int{
				{1},
				{1, 1},
			},
		},
		"Five Rows": {
			numRows: 5,
			expected: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := pascalstriangle.GeneratePascalsTriangle(tc.numRows)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Fatalf("Expected: %v, Got: %v", tc.expected, got)
			}
		})
	}
}

func TestGeneratePascalsTriangleRows(t *testing.T) {
	testCases := map[string]struct {
		rowIndex int
		expected []int
	}{
		"First Row": {
			rowIndex: 0,
			expected: []int{1},
		},
		"Second Row": {
			rowIndex: 1,
			expected: []int{1, 1},
		},
		"Fifth Row": {
			rowIndex: 4,
			expected: []int{1, 4, 6, 4, 1},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := pascalstriangle.GeneratePascalsTriangleRow(tc.rowIndex)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Fatalf("Expected: %v, Got: %v", tc.expected, got)
			}
		})
	}
}
