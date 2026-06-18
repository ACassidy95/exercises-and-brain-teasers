package clockangle_test

import (
	clockangle "daily/clock-angle"
	"testing"
)

func TestClockAngle(t *testing.T) {
	testCases := map[string]struct {
		hour     int
		minute   int
		expected float64
	}{
		"12:30": {
			hour:     12,
			minute:   30,
			expected: 165,
		},
		"3:30": {
			hour:     3,
			minute:   30,
			expected: 75,
		},
		"3:15": {
			hour:     3,
			minute:   15,
			expected: 7.5,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := clockangle.ClockAngle(tc.hour, tc.minute)
			if got != tc.expected {
				t.Fatalf("Expected: %v, Got: %v", tc.expected, got)
			}
		})
	}
}
