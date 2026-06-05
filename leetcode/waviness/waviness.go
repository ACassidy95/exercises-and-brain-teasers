package waviness

import "math"

func TotalWaviness(num1 int, num2 int) int {
	var waviness, min, max int
	var digits []int

	// Only three-digit or greater numbers count, thus if upper-bound
	// is lower, just return 0.
	if num2 < 100 {
		return 0
	}

	// Two digit numbers have waviness 0, so begin at 100
	if num1 < 100 {
		min = 100
	} else {
		min = num1
	}
	max = num2

	// Calculate the nearest power of 10 not greater than
	// the input value
	magnitude := func(x int) int {
		m := 1
		for {
			x /= 10
			if x == 0 {
				break
			}
			m++
		}
		return m
	}

	for i := min; i <= max; i++ {
		m := int(math.Pow(10., float64(magnitude(i))))
		for m > 0 {
			digits = append(digits, i/m)
			i -= m
			m /= 10
		}

		for j := 1; j < len(digits)-2; j++ {
			if digits[j-1] < digits[j] {
				waviness++
			}
			if digits[j] < digits[j+1] {
				waviness++
			}
		}
	}

	return 0
}
