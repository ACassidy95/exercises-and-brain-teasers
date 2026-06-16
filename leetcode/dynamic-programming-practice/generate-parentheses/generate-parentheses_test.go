package generateparentheses_test

import (
	generateparentheses "dynamic-programming-practice/generate-parentheses"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestGenerateParentheses(t *testing.T) {
	testCases := map[string]struct {
		n        int
		expected []string
	}{
		"n=1": {
			n:        1,
			expected: []string{"()"},
		},
		"n=2": {
			n:        2,
			expected: []string{"(())", "()()"},
		},
		"n=3": {
			n:        3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		"n=4": {
			n:        4,
			expected: nil,
		},
		"n=5": {
			n:        5,
			expected: nil,
		},
		"n=6": {
			n:        6,
			expected: nil,
		},
		"n=7": {
			n:        7,
			expected: nil,
		},
		"n=8": {
			n:        8,
			expected: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := generateparentheses.GenerateParentheses(tc.n)
			diffIgnoreOrder := cmp.Diff(got, tc.expected, cmpopts.SortSlices(func(a, b string) bool { return a < b }))
			if diffIgnoreOrder != "" {
				t.Fatalf("Expected: %v, Got: %v, Diff: %v", tc.expected, got, diffIgnoreOrder)
			}
		})
	}
}
