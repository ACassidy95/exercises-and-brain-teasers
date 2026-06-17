package jump

import "math"

func Jump(nums []int) int {
	var dp []int

	// A list of one number takes 0 steps to reach the end
	if len(nums) < 2 {
		return 0
	}

	minNonZero := func(ints ...int) int {
		m := math.MaxInt
		for _, i := range ints {
			if i < m {
				m = i
			}
		}
		return m
	}

	dp = make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] == 0 {
			dp[i] = 0
		} else {
			rang := dp[i:min(len(dp), i+nums[i])]
			dp[i] = dp[i] + 1 + minNonZero(rang...)
		}
	}

	return dp[0]
}
