package climbingstairs_test

import (
	climbingstairs "dynamic-programming-practice/climbing-stairs"
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	testCases := map[string]struct {
		numStairs int
		expected  int
	}{
		"No stairs": {
			numStairs: 0,
			expected:  1,
		},
		"1 stair": {
			numStairs: 1,
			expected:  1,
		},
		"2 stairs": {
			numStairs: 2,
			expected:  2,
		},
		"3 stairs": {
			numStairs: 3,
			expected:  3,
		},
		"45 stairs": {
			numStairs: climbingstairs.MAX_STAIRS,
			expected:  1836311903,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := climbingstairs.ClimbStairs(tc.numStairs)
			if got != tc.expected {
				t.Fatalf("Expected %v, got %v", tc.expected, got)
			}
		})
	}

	var total uint64
	for i, count := range climbingstairs.ClimbStairsCalls {
		fmt.Printf("ClimbStairs(%d) call count: %d\n", i, count)
		total += count
	}
	fmt.Printf("Total calls to ClimbStairs: %d\n", total)
}

func TestClimbStairsMemo(t *testing.T) {
	testCases := map[string]struct {
		numStairs int
		expected  int
	}{
		"No stairs": {
			numStairs: 0,
			expected:  1,
		},
		"1 stair": {
			numStairs: 1,
			expected:  1,
		},
		"2 stairs": {
			numStairs: 2,
			expected:  2,
		},
		"3 stairs": {
			numStairs: 3,
			expected:  3,
		},
		"45 stairs": {
			numStairs: climbingstairs.MAX_STAIRS,
			expected:  1836311903,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			memo := make([]int, climbingstairs.MAX_STAIRS+1)
			got := climbingstairs.ClimbStairsMemo(tc.numStairs, memo)
			if got != tc.expected {
				t.Fatalf("Expected %v, got %v", tc.expected, got)
			}
		})
	}

	var total uint64
	for i, count := range climbingstairs.ClimbStairsMemoCalls {
		fmt.Printf("ClimbStairsMemo(%d) call count: %d\n", i, count)
		total += count
	}
	fmt.Printf("Total calls to ClimbStairsMemo: %d\n", total)
}

func TestClimbStairsTab(t *testing.T) {
	testCases := map[string]struct {
		numStairs int
		expected  int
	}{
		"No stairs": {
			numStairs: 0,
			expected:  1,
		},
		"1 stair": {
			numStairs: 1,
			expected:  1,
		},
		"2 stairs": {
			numStairs: 2,
			expected:  2,
		},
		"3 stairs": {
			numStairs: 3,
			expected:  3,
		},
		"45 stairs": {
			numStairs: climbingstairs.MAX_STAIRS,
			expected:  1836311903,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := climbingstairs.ClimbStairsTab(tc.numStairs)
			if got != tc.expected {
				t.Fatalf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}
