package jump

import "math"

func Jump(nums []int) int {
	var dp []int
	var l int

	l = len(nums)

	// A list of one number takes 0 steps to reach the end
	// and a list of two takes 1 (given we are guaranteed to reach the end)
	switch l {
	case 1:
		return 0
	case 2:
		return 1
	}

	minNonZero := func(ints ...int) int {
		m := math.MaxInt
		for _, i := range ints {
			if i < m && i != 0 {
				m = i
			}
		}

		if m == math.MaxInt {
			m = 0
		}
		return m
	}

	/*
	 * dp[i] := the minimum number of steps to reach the final position
	 * If the current number is 0, we cannot move from that position and there are no routes to the final position i.e. dp[i] = 0
	 * If the current number is big enough to move to or over the final position, it takes a single step, i.e. dp[i] = 1
	 * If the current number is not big enough to move to the final position, it takes 1 + min of the non-zero steps in range of the current number,
	 * i.e. dp[i] = 1 + min(dp[i:i+nums[i]])
	 * (0s are dead ends)
	 */
	dp = make([]int, l)
	for i := l - 2; i >= 0; i-- {
		curr := nums[i]
		if curr == 0 {
			dp[i] = 0
		} else if curr >= l-1-i {
			dp[i] = 1
		} else {
			mRange := dp[i+1 : i+curr+1]
			m := minNonZero(mRange...)
			if m == 0 {
				dp[i] = 0
			} else {
				dp[i] = dp[i] + 1 + m
			}
		}
	}

	return dp[0]
}
