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
