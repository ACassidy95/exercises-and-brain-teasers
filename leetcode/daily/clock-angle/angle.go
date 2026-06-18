package clockangle

func ClockAngle(hour, minutes int) float64 {
	// minStep: Angle made between one minute and the next
	// hrStep: Angle made between one hour and the next
	// theta_min: Angle from 12 (0 degrees) to minute value
	// theta_hr: Angle from 12 (0 degrees) to hour value
	const minStep, hrStep float64 = 6, 30
	var theta_min, theta_hr float64

	theta_min = float64(minutes) * minStep

	theta_hr = float64(hour%12) * hrStep
	theta_hr += hrStep * (theta_min / 360)

	absDiff := max(theta_hr, theta_min) - min(theta_hr, theta_min)
	if absDiff <= 180 {
		return absDiff
	} else {
		return 360 - absDiff
	}
}
