package setmismatch_test

import (
	setmismatch "quest/set-mismatch"
	"reflect"
	"testing"
)

func TestSetMismatch(t *testing.T) {
	testCases := map[string]struct {
		nums     []int
		expected []int
	}{
		"[1, 2, 2, 4]": {
			nums:     []int{1, 2, 2, 4},
			expected: []int{2, 3},
		},
		"[1, 1]": {
			nums:     []int{1, 1},
			expected: []int{1, 2},
		},
		"[2, 2]": {
			nums:     []int{2, 2},
			expected: []int{2, 1},
		},
		"[3, 2, 2]": {
			nums:     []int{3, 2, 2},
			expected: []int{2, 1},
		},
		"[2, 3, 2]": {
			nums:     []int{2, 3, 2},
			expected: []int{2, 1},
		},
		"[1, 2, 3, 4, 5, 5, 7, 8, 9]": {
			nums:     []int{1, 2, 3, 4, 5, 5, 7, 8, 9},
			expected: []int{5, 6},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := setmismatch.SetMismatch(tc.nums)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Fatalf("Expected: %v, Got: %v", tc.expected, got)
			}
		})
	}
}
