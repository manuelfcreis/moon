package moon

import (
	"math"
	"time"
)

const (
	daysBetweenFullMoons = 29.53
)

// dates can't be defined as a const :(
func originalFullMoon() time.Time {
	return time.Date(2024, 05, 23, 0, 0, 0, 0, time.UTC)
}

func DaysToFullMoon() time.Time {
	daysSinceFullMoon := time.Now().Sub(originalFullMoon()).Hours() / 24
	daysUntilNextFullMoon := time.Duration(daysBetweenFullMoons - daysSinceFullMoon*24)
	targetTime := time.Now().Add(time.Hour * daysUntilNextFullMoon)

	return targetTime
}

func calculateMoonPercentage(currentDate time.Time) float64 {
	daysSinceFullMoon := math.Round(currentDate.Sub(originalFullMoon()).Hours() / 24)
	modulos := math.Mod(daysSinceFullMoon, daysBetweenFullMoons)
	moonPercentage := modulos / daysBetweenFullMoons * 8

	return moonPercentage
}

func Moon(targetTime time.Time) (string, string) {
	moonPhasePercentage := calculateMoonPercentage(targetTime)
	moonInformation := information(moonPhasePercentage)
	moonDrawing := drawMoon(moonPhasePercentage)

	return moonInformation, moonDrawing
}
