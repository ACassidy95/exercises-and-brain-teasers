// Given the array nums, for each nums[i] find out how many numbers in the array are smaller than it.
// That is, for each nums[i] you have to count the number of valid j's such that j != i and nums[j] < nums[i].

// Return the answer in an array.

package smallernumbers

func SmallerNumbersThanCurrent(nums []int) []int {
	var counts []int

	counts = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] < nums[i] {
				counts[i]++
			}
		}
	}

	return counts
}

func SmallerNumbersThanCurrentOpt(nums []int) []int {
	const maxNumsLen int = 101
	var dp, counts []int

	dp = make([]int, maxNumsLen)
	counts = make([]int, len(nums))
	for _, num := range nums {
		dp[num]++
	}

	for i := 1; i < maxNumsLen-1; i++ {
		dp[i] = dp[i] + dp[i-1]
	}

	for i, num := range nums {
		counts[i] = dp[num]
	}

	return counts
}
