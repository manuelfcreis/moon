package moon

import "math"

func determineMoonPhase(moonPhasePercentage float64) string {
  var moonPhase string

  switch int64(math.Round(moonPhasePercentage)) {
  case 0: moonPhase = "Full moon"
  case 1: moonPhase = "Waning gibbous moon"
  case 2: moonPhase = "Last quarter moon"
  case 3: moonPhase = "Waning crescent moon"
  case 4: moonPhase = "New moon"
  case 5: moonPhase = "Waxing crescent moon"
  case 6: moonPhase = "First quarter moon"
  case 7: moonPhase = "Waxing gibbous moon"
  case 8: moonPhase = "Full moon"
  }

  return moonPhase
}

