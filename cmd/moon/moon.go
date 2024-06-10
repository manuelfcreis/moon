package moon

import "time"

const (
  daysBetweenFullMoons = 29.53
)

// dates can't be defined as a const :(
func originalFullMoon() time.Time {
  return time.Date(2024, 05, 23, 0, 0, 0, 0, time.UTC)
}

func DaysToFullMoon() time.Time {
  daysSinceFullMoon := time.Now().Sub(originalFullMoon()).Hours() / 24
  daysUntilNextFullMoon := time.Duration(daysBetweenFullMoons - daysSinceFullMoon * 24)
  targetTime := time.Now().Add(time.Hour * daysUntilNextFullMoon)

  return targetTime
}

func calculateMoonPercentage(currentDate time.Time) float64 {
  daysSinceFullMoon := currentDate.Sub(originalFullMoon()).Hours() / 24

  return daysSinceFullMoon / daysBetweenFullMoons * 8
}


func Moon(targetTime time.Time) (string, string) {
  moonPhasePercentage := calculateMoonPercentage(targetTime)
  moonPhase := determineMoonPhase(moonPhasePercentage)
  moonDrawing := drawMoon(moonPhasePercentage)

  return moonPhase, moonDrawing
}
