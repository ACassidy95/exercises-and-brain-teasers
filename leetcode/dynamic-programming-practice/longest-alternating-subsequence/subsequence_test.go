package longestalternatingsubsequence_test

import (
	longestalternatingsubsequence "dynamic-programming-practice/longest-alternating-subsequence"
	"reflect"
	"testing"
)

func TestLongestAlternatingSubsequence(t *testing.T) {
	testCases := map[string]struct {
		words    []string
		groups   []int
		expected []string
	}{
		"[e, a, b]/[0, 0, 1]": {
			words:    []string{"e", "a", "b"},
			groups:   []int{0, 0, 1},
			expected: []string{"e", "a"},
		},
		"[a, b, c, d]/[1, 0, 1, 1]": {
			words:    []string{"a", "b", "c", "d"},
			groups:   []int{1, 0, 1, 1},
			expected: []string{"a", "b", "c"},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := longestalternatingsubsequence.LongestAlternatingSubsequence(tc.words, tc.groups)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Fatalf("Expected: %v, Got: %v", tc.expected, got)
			}
		})
	}
}
