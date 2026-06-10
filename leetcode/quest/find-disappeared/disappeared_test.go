package finddisappeared_test

import (
	finddisappeared "quest/find-disappeared"
	"reflect"
	"testing"
)

func TestFindDisappeared(t *testing.T) {
	testCases := map[string]struct {
		nums     []int
		expected []int
	}{
		"[4,3,2,7,8,2,3,1]": {
			nums:     []int{4, 3, 2, 7, 8, 2, 3, 1},
			expected: []int{5, 6},
		},
		"[1,1]": {
			nums:     []int{1, 1},
			expected: []int{2},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := finddisappeared.FindDisappeared(tc.nums)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Fatalf("Expected: %v, Got: %v", tc.expected, got)
			}
		})
	}
}
