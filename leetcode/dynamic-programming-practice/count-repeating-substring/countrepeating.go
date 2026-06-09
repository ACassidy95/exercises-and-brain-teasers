// For a string sequence, a string word is k-repeating if word
// concatenated k times is a substring of sequence.
// The word's maximum k-repeating value is the highest value k
// where word is k-repeating in sequence.
// If word is not a substring of sequence, word's maximum k-repeating value is 0.

// Given strings sequence and word, return the maximum k-repeating
// value of word in sequence.

package countrepeatingsubstring

func CountRepeatingSubstring(sequence, word string) int {
	var dp []int
	var dpMax int

	slen := len(sequence)
	wlen := len(word)

	dp = make([]int, slen)

	for i := slen - 1; i >= wlen-1; i-- {
		var j, k int
		for j, k = wlen-1, i; j >= 0 && word[j] == sequence[k]; j, k = j-1, k-1 {
		}

		if k == i-wlen {
			if i >= slen-wlen {
				dp[i] = dp[k+wlen] + 1
			} else {
				dp[i] = dp[i+wlen] + 1
			}
		}
	}

	for _, d := range dp {
		if d > dpMax {
			dpMax = d
		}
	}

	return dpMax
}
