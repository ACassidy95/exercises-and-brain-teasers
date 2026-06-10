package setmismatch

func SetMismatch(nums []int) []int {
	var xorAll, xorNums, xorRes, rightBit, xorSet, xorUnset int

	for i, num := range nums {
		xorAll ^= i + 1
		xorNums ^= num
	}

	xorRes = xorAll ^ xorNums
	rightBit = xorRes & -xorRes

	for i, num := range nums {
		if (i+1)&rightBit != 0 {
			xorSet ^= (i + 1)
		} else {
			xorUnset ^= (i + 1)
		}

		if num&rightBit != 0 {
			xorSet ^= num
		} else {
			xorUnset ^= num
		}
	}

	for _, num := range nums {
		if xorSet == num {
			return []int{xorSet, xorUnset}
		}
	}
	return []int{xorUnset, xorSet}
}
