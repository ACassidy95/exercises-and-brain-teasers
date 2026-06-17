package jump

import "math"

func Jump(nums []int) int {
	var dp []int

	// A list of one number takes 0 steps to reach the end
	// and a list of two takes 1 (given we are guaranteed to reach the end)
	if len(nums) == 1 {
		return 0
	} else if len(nums) == 2 {
		return 1
	}

	minNonZero := func(ints ...int) int {
		m := math.MaxInt
		for _, i := range ints {
			if i < m {
				m = i
			}
		}

		if m == math.MaxInt {
			m = 0
		}
		return max(1, m)
	}

	/*
	 * dp[i] := the minimum number of steps to reach the final position
	 * If the current number is 0, we cannot move from that position and there are no routes to the final position i.e. dp[i] = 0
	 * If the current number is big enough to move to or over the final position, it takes a single step, i.e. dp[i] = 1
	 * If the current number is not big enough to move to the final position, it takes 1 + min of the non-zero steps in range of the current number,
	 * i.e. dp[i] = 1 + min(dp[i:i+nums[i]])
	 * (0s are dead ends)
	 */
	dp = make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] == 0 {
			dp[i] = 0
		} else if nums[i] >= len(nums)-1-i {
			dp[i] = 1
		} else {
			m := min(len(dp), i+nums[i])
			mm := m + nums[i]
			mRange := dp[m:mm]
			dp[i] = dp[i] + 1 + minNonZero(mRange...)
		}
	}

	return dp[0]
}
