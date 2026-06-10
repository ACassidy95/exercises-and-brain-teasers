package finddisappeared

func FindDisappeared(nums []int) []int {
	var freqs, missing []int

	freqs = make([]int, len(nums)+1)
	for _, num := range nums {
		freqs[num]++
	}

	for i, freq := range freqs {
		if freq == 0 && i > 0 {
			missing = append(missing, i)
		}
	}

	return missing
}
