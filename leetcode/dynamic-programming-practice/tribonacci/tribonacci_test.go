package tribonacci_test

import (
	"dynamic-programming-practice/tribonacci"
	"testing"
)

func TestTribonacci(t *testing.T) {
	testCases := map[string]struct {
		n        int
		expected int
	}{
		"T(4)": {
			n:        4,
			expected: 4,
		},
		"T(25)": {
			n:        25,
			expected: 1389537,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tribonacci.Tribonacci(tc.n)
			if got != tc.expected {
				t.Fatalf("Expected: %v, Got: %v", tc.expected, got)
			}
		})
	}
}
