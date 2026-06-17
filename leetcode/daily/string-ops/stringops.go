package stringops

import (
	"slices"
)

func ProcessStringOps(s string) string {
	var chars []rune

	for _, char := range s {
		switch char {
		case '*':
			if len(chars) > 1 {
				chars = chars[:len(chars)-1]
			} else {
				chars = []rune{}
			}
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

func ProcessSpecialStringOps(s string, k int64) byte {
	var totalLen int64

	for _, char := range s {
		switch char {
		case '*':
			if totalLen >= 1 {
				totalLen--
			}
		case '#':
			totalLen *= 2
		case '%':
			// Continue here to avoid adding where there is no need to
			continue
		default:
			totalLen++
		}
	}

	if k >= totalLen {
		return '.'
	}

	r := []rune(s)
	for i := len(r) - 1; i >= 0; i-- {
		switch r[i] {
		case '*':
			totalLen++
		case '#':
			halfLen := totalLen / 2
			if k >= halfLen {
				k -= halfLen
			}
			totalLen = halfLen
		case '%':
			k = totalLen - 1 - k
		default:
			if k == totalLen-1 {
				return byte(r[i])
			}
			totalLen--
		}
	}

	return '.'
}
