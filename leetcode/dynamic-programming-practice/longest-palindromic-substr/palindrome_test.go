package longestpalindromicsubstr_test

import (
	longestpalindromicsubstr "dynamic-programming-practice/longest-palindromic-substr"
	"testing"
)

func TestLongestPaalindromicSubstring(t *testing.T) {
	testCases := map[string]struct {
		s        string
		expected string
	}{
		"babad -> bab": {
			s:        "babad",
			expected: "bab",
		},
		"cbbd -> bb": {
			s:        "cbbd",
			expected: "bb",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := longestpalindromicsubstr.LongestPalindromicSubstr(tc.s)
			if tc.expected != got {
				t.Fatalf("Expected: %v, Got: %v", tc.expected, got)
			}
		})
	}
}
