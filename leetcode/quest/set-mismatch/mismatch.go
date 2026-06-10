package setmismatch

func SetMismatch(nums []int) []int {
	var dup, xor int

	xor ^= nums[0]
	for i := 1; i < len(nums); i++ {
		xor ^= nums[i]
		if nums[i] == nums[i-1] {
			dup = nums[i]
		}
	}

	xor ^= dup
	for i := 1; i <= len(nums); i++ {
		xor ^= i
	}

	return []int{dup, xor}
}
