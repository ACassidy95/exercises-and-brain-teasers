package clockangle

func ClockAngle(hour, minute int) float64 {
	// minStep: Angle made between one minute and the next
	// hrStep: Angle made between one hour and the next
	// theta_min: Angle from 12 (0 degrees) to minute value
	// theta_hr: Angle from 12 (0 degrees) to hour value
	// theta_offset: Angle between the line 12-6 and the line formed by the minute hand
	// theta_oh: Angle between the offset and the hour
	// theta_mh: Angle between the minute and the hour
	const minStep, hrStep float64 = 6, 30
	var theta_min, theta_hr, theta_offset, theta_oh, theta_mh float64

	mod := func(a, b int) int {
		r := a % b
		if r < 0 {
			return r + b
		} else {
			return r
		}
	}

	theta_min = float64(minute) * minStep

	theta_hr = float64(hour%12) * hrStep
	theta_hr += hrStep * (theta_min / 360)

	theta_offset = float64(mod(minute-30, 12)) * minStep

	theta_oh = theta_hr - theta_offset
	theta_mh = 180 - theta_oh
	return theta_mh
}
