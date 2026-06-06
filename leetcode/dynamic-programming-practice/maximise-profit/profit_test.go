package maximiseprofit_test

import (
	maximiseprofit "dynamic-programming-practice/maximise-profit"
	"testing"
)

func TestMaximiseProfit(t *testing.T) {
	testCases := map[string]struct {
		prices   []int
		expected int
	}{
		"Profitable": {
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		"Unprofitable": {
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := maximiseprofit.MaximiseProfit(tc.prices)
			if got != tc.expected {
				t.Fatalf("Expected: %v, Got: %v", tc.expected, got)
			}
		})
	}
}
