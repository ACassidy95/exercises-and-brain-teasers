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
