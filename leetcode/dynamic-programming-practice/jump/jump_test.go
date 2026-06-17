package jump_test

import (
	"dynamic-programming-practice/jump"
	"testing"
)

func TestJump(t *testing.T) {
	testCases := map[string]struct {
		nums     []int
		expected int
	}{
		"[2,3,1,1,4]": {
			nums:     []int{2, 3, 1, 1, 4},
			expected: 2,
		},
		"[2,3,0,1,4]": {
			nums:     []int{2, 3, 0, 1, 4},
			expected: 2,
		},
		"[1,2]": {
			nums:     []int{1, 2},
			expected: 1,
		},
		"[3,2,1]": {
			nums:     []int{3, 2, 1},
			expected: 1,
		},
		"[1,1,1,1]": {
			nums:     []int{1, 1, 1, 1},
			expected: 3,
		},
		"[2,1,1,1,1]": {
			nums:     []int{2, 1, 1, 1, 1},
			expected: 3,
		},
		"[2,0,2,4,6,0,0,3]": {
			nums:     []int{2, 0, 2, 4, 6, 0, 0, 3},
			expected: 3,
		},
		"[8,6,5,2,1,8,1,8,9,7,1,9,1,0,0,3,2,3,5,8,9,4,3,6,5,9,7,9,9,7,3,0,5,1,4,8,9]": {
			nums:     []int{8, 6, 5, 2, 1, 8, 1, 8, 9, 7, 1, 9, 1, 0, 0, 3, 2, 3, 5, 8, 9, 4, 3, 6, 5, 9, 7, 9, 9, 7, 3, 0, 5, 1, 4, 8, 9},
			expected: 5,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := jump.JumpOpt(tc.nums)
			if got != tc.expected {
				t.Fatalf("Expected: %v, Got: %v", tc.expected, got)
			}
		})
	}
}
