// Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n),
// ans[i] is the number of 1's in the binary representation of i.

package countingbits

func CountBits(n int) []int {
	var bitcounts []int

	bitcounts = append(bitcounts, 0)
	for i := 1; i <= n; i++ {

		// Powers of two only contain 1 zero so can skip computation
		if (i & (i - 1)) == 0 {
			bitcounts = append(bitcounts, 1)
		} else {
			if i%2 != 0 {
				bitcounts = append(bitcounts, bitcounts[i-1]+1)
			} else {
				bitcounts = append(bitcounts, bitcounts[i-1])
			}
		}
	}

	return bitcounts
}
