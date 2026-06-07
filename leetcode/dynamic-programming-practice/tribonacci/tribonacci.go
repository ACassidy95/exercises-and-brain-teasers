// T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.
// Given n, return the value of Tn.

package tribonacci

func Tribonacci(n int) int {
	const maxN int = 37
	var dp []int

	dp = make([]int, maxN+1)

	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= maxN; i++ {
		dp[i] = dp[i-3] + dp[i-2] + dp[i-1]
	}

	return dp[n]
}
