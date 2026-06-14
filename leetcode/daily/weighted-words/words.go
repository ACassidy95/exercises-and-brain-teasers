package weightedwords

import "strings"

func WeightedWords(words []string, weights []int) string {
	var sb strings.Builder

	for _, word := range words {
		var wweight int
		var cmap rune

		for _, char := range word {
			wweight += weights[int(char-'a')]
		}

		wweight %= 26
		// cmap = rune(int('z') - wweight)
		cmap = 'z' - rune(wweight)
		sb.WriteRune(cmap)
	}

	return sb.String()
}
