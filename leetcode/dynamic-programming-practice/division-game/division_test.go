package divisiongame_test

import (
	divisiongame "dynamic-programming-practice/division-game"
	"testing"
)

func TestDivisionGame(t *testing.T) {
	testCases := map[string]struct {
		n        int
		expected bool
	}{
		"2": {
			n:        2,
			expected: true,
		},
		"3": {
			n:        3,
			expected: false,
		},
		"4": {
			n:        4,
			expected: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := divisiongame.DivisionGame(tc.n)
			if tc.expected != got {
				t.Fatalf("Expected: %v, Got: %v", tc.expected, got)
			}
		})
	}
}
