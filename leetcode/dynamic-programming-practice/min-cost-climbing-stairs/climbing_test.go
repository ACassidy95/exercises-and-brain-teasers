package mincostclimbingstairs_test

import (
	mincostclimbingstairs "dynamic-programming-practice/min-cost-climbing-stairs"
	"testing"
)

func TestClimbStairsWithCost(t *testing.T) {
	testCases := map[string]struct {
		cost     []int
		expected int
	}{
		"Staircase 1": {
			cost:     []int{10, 15, 20},
			expected: 15,
		},
		"Staircase 2": {
			cost:     []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			expected: 6,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := mincostclimbingstairs.ClimbStairsWithCost(tc.cost)
			if got != tc.expected {
				t.Fatalf("Expected: %v, Got: %v", tc.expected, got)
			}
		})
	}
}
