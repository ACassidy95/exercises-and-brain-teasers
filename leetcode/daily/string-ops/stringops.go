package stringops

import "slices"

func ProcessStringOps(s string) string {
	var chars []rune

	for _, char := range s {
		switch char {
		case '*':
			chars = chars[:len(chars)-1]
		case '#':
			chars = append(chars, chars...)
		case '%':
			slices.Reverse(chars)
		default:
			chars = append(chars, char)
		}
	}

	return string(chars)
}
