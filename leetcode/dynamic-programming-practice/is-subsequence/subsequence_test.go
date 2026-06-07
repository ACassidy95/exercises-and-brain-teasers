package issubsequence_test

import (
	issubsequence "dynamic-programming-practice/is-subsequence"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	testCases := map[string]struct {
		s        string
		t        string
		expected bool
	}{
		"Same Sequence": {
			s:        "helloworld",
			t:        "helloworld",
			expected: true,
		},
		"Continuous Subsequence": {
			s:        "low",
			t:        "helloworld",
			expected: true,
		},
		"Discontinuous Subsequence": {
			s:        "abc",
			t:        "ahbgdc",
			expected: true,
		},
		"No subsequence": {
			s:        "axc",
			t:        "ahbgdc",
			expected: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := issubsequence.IsSubsequence(tc.s, tc.t)
			if got != tc.expected {
				t.Fatalf("Expected: %v, Got: %v", tc.expected, got)
			}
		})
	}
}
