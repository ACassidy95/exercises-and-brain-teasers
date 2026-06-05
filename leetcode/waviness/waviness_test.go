package waviness_test

import (
	"testing"
	"waviness"
)

func TestWaviness(t *testing.T) {
	tests := map[string]struct {
		lower    int
		upper    int
		expected int
	}{
		"Single number range less than 100": {
			lower:    99,
			upper:    99,
			expected: 0,
		},
		"Multi-number range less than 100": {
			lower:    1,
			upper:    99,
			expected: 0,
		},
		"99-101": {
			lower:    99,
			upper:    101,
			expected: 1,
		},
		"120-130": {
			lower:    120,
			upper:    130,
			expected: 3,
		},
		"198-202": {
			lower:    198,
			upper:    202,
			expected: 3,
		},
		"4848-4848": {
			lower:    4848,
			upper:    4848,
			expected: 2,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := waviness.TotalWaviness(tc.lower, tc.upper)
			if got != tc.expected {
				t.Fatalf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}
