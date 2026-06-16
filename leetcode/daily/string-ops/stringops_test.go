package stringops_test

import (
	stringops "daily/string-ops"
	"testing"
)

func TestProcessStringOps(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected string
	}{
		"a#b%*": {
			input:    "a#b%*",
			expected: "ba",
		},
		"z*#": {
			input:    "z*#",
			expected: "",
		},
		"*%": {
			input:    "*%",
			expected: "",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := stringops.ProcessStringOps(tc.input)
			if got != tc.expected {
				t.Fatalf("Expected: %v, Got: %v", tc.expected, got)
			}
		})
	}
}
