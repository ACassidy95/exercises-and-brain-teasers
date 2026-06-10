package setmismatch_test

import (
	setmismatch "quest/set-mismatch"
	"reflect"
	"testing"
)

func TestSetMismatch(t *testing.T) {
	testCases := map[string]struct {
		nums     []int
		expected []int
	}{}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := setmismatch.SetMismatch(tc.nums)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Fatalf("Expected: %v, Got: %v", tc.expected, got)
			}
		})
	}
}
