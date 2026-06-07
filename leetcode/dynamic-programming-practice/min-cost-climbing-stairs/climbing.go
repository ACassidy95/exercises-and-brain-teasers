// You are given an integer array cost where cost[i] is the cost of ith step on a staircase.
// Once you pay the cost, you can either climb one or two steps.
// You can either start from the step with index 0, or the step with index 1.
// Return the minimum cost to reach the top of the floor.

package mincostclimbingstairs

func ClimbStairsWithCost(cost []int) int {
	// dp will track the minimum cost to reach the top from position i
	var dp []int
	var min func(int, int) int

	min = func(m, n int) int {
		if m < n {
			return m
		} else {
			return n
		}
	}

	dp = make([]int, len(cost)+1)

	for i := len(cost) - 1; i >= 0; i-- {
		if i+1 >= len(cost) {
			dp[i] = cost[i]
		} else {
			dp[i] = cost[i] + min(dp[i+1], dp[i+2])
		}

	}

	return min(dp[0], dp[1])
}
