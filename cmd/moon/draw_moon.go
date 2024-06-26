package moon

import (
	"math"
	"strings"
)

func fillMoon(moonPhasePercentage float64, row []string, negative string) {
	positive := ":"
	moonPhasePercentage = math.Round(moonPhasePercentage)
	afterNewMoon := moonPhasePercentage >= 4

	if afterNewMoon {
		moonPhasePercentage -= 4
	}

	for i := range row {
		rowPercentage := float64(i) / float64(len(row)) * 4

		if rowPercentage > (4 - moonPhasePercentage) {
			if afterNewMoon {
				row[i] = positive
			} else {
				row[i] = negative
			}
		} else {
			if afterNewMoon {
				row[i] = negative
			} else {
				row[i] = positive
			}
		}
	}
}

func drawMoon(moonPhasePercentage float64) string {
	var b strings.Builder
	var topRow [7]string
	var bottomRow [3]string
	var firstMoonRow [9]string
	var secondMoonRow [9]string
	var thirdMoonRow [9]string

	fillMoon(moonPhasePercentage, topRow[:], " ")
	fillMoon(moonPhasePercentage, firstMoonRow[:], " ")
	fillMoon(moonPhasePercentage, secondMoonRow[:], " ")
	fillMoon(moonPhasePercentage, thirdMoonRow[:], " ")
	fillMoon(moonPhasePercentage, bottomRow[:], ".")

	if thirdMoonRow[0] == " " {
		thirdMoonRow[0] = "." // This is just a nudge to make the New moon prettier
	}

	b.WriteString("       _..._\n")
	b.WriteString("     ." + strings.Join(topRow[:], "") + ".\n")
	b.WriteString("    :" + strings.Join(firstMoonRow[:], "") + ":\n")
	b.WriteString("    :" + strings.Join(secondMoonRow[:], "") + ":\n")
	b.WriteString("    '" + strings.Join(thirdMoonRow[:], "") + "'\n")
	b.WriteString("      ''" + strings.Join(bottomRow[:], "") + "''\n")

	return b.String()
}
