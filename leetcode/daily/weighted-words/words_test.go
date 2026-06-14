package weightedwords_test

import (
	weightedwords "daily/weighted-words"
	"testing"
)

func TestWeightedWords(t *testing.T) {
	testCases := map[string]struct {
		words    []string
		weights  []int
		expected string
	}{
		"[\"abcd\",\"def\",\"xyz\"]": {
			words:    []string{"abcd", "def", "xyz"},
			weights:  []int{5, 3, 12, 14, 1, 2, 3, 2, 10, 6, 6, 9, 7, 8, 7, 10, 8, 9, 6, 9, 9, 8, 3, 7, 7, 2},
			expected: "rij",
		},
		"[\"a\",\"b\",\"c\"]": {
			words:    []string{"a", "b", "c"},
			weights:  []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			expected: "yyy",
		},
		"[\"abcd\"]": {
			words:    []string{"abcd"},
			weights:  []int{7, 5, 3, 4, 3, 5, 4, 9, 4, 2, 2, 7, 10, 2, 5, 10, 6, 1, 2, 2, 4, 1, 3, 4, 4, 5},
			expected: "g",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := weightedwords.WeightedWords(tc.words, tc.weights)
			if got != tc.expected {
				t.Fatalf("Expected: %v, Got: %v", tc.expected, got)
			}
		})
	}
}
