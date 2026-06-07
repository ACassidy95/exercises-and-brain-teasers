//Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
// A subsequence of a string is a new string that is formed from the original
// string by deleting some (can be none) of the characters without disturbing
// the relative positions of the remaining characters.
// (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

package issubsequence

func IsSubsequence(s, t string) bool {
	var found bool
	var i, j int

	if s == "" {
		return true
	}

	for i, j = 0, 0; i < len(s) && j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}
	}

	if i == len(s) {
		found = true
	}

	return found
}
