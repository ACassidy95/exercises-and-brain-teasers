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
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := jump.Jump(tc.nums)
			if got != tc.expected {
				t.Fatalf("Expected: %v, Got: %v", tc.expected, got)
			}
		})
	}
}
