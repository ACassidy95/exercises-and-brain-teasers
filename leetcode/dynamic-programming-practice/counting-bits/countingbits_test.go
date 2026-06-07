package countingbits_test

import (
	countingbits "dynamic-programming-practice/counting-bits"
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	testCases := map[string]struct {
		n        int
		expected []int
	}{
		"Zero": {
			n:        0,
			expected: []int{0},
		},
		"One": {
			n:        1,
			expected: []int{0, 1},
		},
		"Two": {
			n:        2,
			expected: []int{0, 1, 1},
		},
		"Five": {
			n:        5,
			expected: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := countingbits.CountBits(tc.n)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Fatalf("Expected %v, Got %v", tc.expected, got)
			}
		})
	}
}
