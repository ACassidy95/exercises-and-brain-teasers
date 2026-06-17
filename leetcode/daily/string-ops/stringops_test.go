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

func TestProcessSpecialStringOps(t *testing.T) {
	testCases := map[string]struct {
		str      string
		k        int64
		expected byte
	}{
		"a#b%*": {
			str:      "a#b%*",
			k:        1,
			expected: 'a',
		},
		"cd%#*#": {
			str:      "cd%#*#",
			k:        3,
			expected: 'd',
		},
		"z*#": {
			str:      "z*#",
			k:        0,
			expected: '.',
		},
		"*%": {
			str:      "*%",
			k:        0,
			expected: '.',
		},
		"%#*gm#xib": {
			str:      "%#*gm#xib",
			k:        2,
			expected: 'g',
		},
		"%%q*#%m##": {
			str:      "%%q*#%m##",
			k:        1,
			expected: 'm',
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := stringops.ProcessSpecialStringOps(tc.str, tc.k)
			if got != tc.expected {
				t.Fatalf("Expected: %v, Got: %v", tc.expected, got)
			}
		})
	}
}
