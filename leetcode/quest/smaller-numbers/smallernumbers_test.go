package smallernumbers_test

import (
	smallernumbers "quest/smaller-numbers"
	"reflect"
	"testing"
)

func TestSmallerNumbersThanCurrent(t *testing.T) {
	testCases := map[string]struct {
		nums     []int
		expected []int
	}{
		"[8,1,2,2,3]": {
			nums:     []int{8, 1, 2, 2, 3},
			expected: []int{4, 0, 1, 1, 3},
		},
		"[6,5,4,8]": {
			nums:     []int{6, 5, 4, 8},
			expected: []int{2, 1, 0, 3},
		},
		"[7,7,7,7]": {
			nums:     []int{7, 7, 7, 7},
			expected: []int{0, 0, 0, 0},
		},
		"[7]": {
			nums:     []int{7},
			expected: []int{0},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := smallernumbers.SmallerNumbersThanCurrent(tc.nums)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Fatalf("Expected: %v, Got: %v", tc.expected, got)
			}
		})
	}
}
