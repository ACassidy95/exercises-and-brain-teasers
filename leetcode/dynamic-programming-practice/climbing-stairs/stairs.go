// You are climbing a staircase. It takes n steps to reach the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

package climbingstairs

const MAX_STAIRS int = 45

var ClimbStairsCalls []uint64 = make([]uint64, MAX_STAIRS+1)
var ClimbStairsMemoCalls []uint64 = make([]uint64, MAX_STAIRS+1)

func ClimbStairs(n int) int {
	ClimbStairsCalls[n]++
	if n == 0 || n == 1 {
		return 1
	}
	return ClimbStairs(n-1) + ClimbStairs(n-2)
}

func ClimbStairsMemo(n int, memo []int) int {
	ClimbStairsMemoCalls[n]++
	if n == 0 || n == 1 {
		memo[n] = 1
		return memo[n]
	}

	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = ClimbStairsMemo(n-1, memo) + ClimbStairsMemo(n-2, memo)
	return memo[n]
}

func ClimbStairsTab(n int) int {
	tab := make([]int, MAX_STAIRS+1)

	tab[0] = 1
	for i := 0; i <= MAX_STAIRS; i++ {
		if (i - 1) >= 0 {
			tab[i] += tab[i-1]
		}
		if (i - 2) >= 0 {
			tab[i] += tab[i-2]
		}
	}

	return tab[n]
}
