package countrepeatingsubstring_test

import (
	countrepeatingsubstring "dynamic-programming-practice/count-repeating-substring"
	"testing"
)

func TestCountRepeatingSubstring(t *testing.T) {
	testCases := map[string]struct {
		sequence string
		word     string
		expected int
	}{
		"sequence = \"ababc\", word = \"ab\"": {
			sequence: "ababc",
			word:     "ab",
			expected: 2,
		},
		"sequence = \"ababc\", word = \"ba\"": {
			sequence: "ababc",
			word:     "ba",
			expected: 1,
		},
		"sequence = \"ababc\", word = \"ac\"": {
			sequence: "ababc",
			word:     "ac",
			expected: 0,
		},
		"sequence = \"ababc\", word = \"\"": {
			sequence: "ababc",
			word:     "ac",
			expected: 0,
		},
		"sequence = \"a\", word = \"a\"": {
			sequence: "a",
			word:     "a",
			expected: 1,
		},
		"sequence = \"aaabaaaabaaabaaaabaaaabaaaabaaaaba\", word = \"aaaba\"": {
			sequence: "aaabaaaabaaabaaaabaaaabaaaabaaaaba",
			word:     "aaaba",
			expected: 5,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := countrepeatingsubstring.CountRepeatingSubstring(tc.sequence, tc.word)
			if got != tc.expected {
				t.Fatalf("Expected: %v, Got: %v", tc.expected, got)
			}
		})
	}
}
