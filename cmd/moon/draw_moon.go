package moon

import "strings"

func fillMoon(moonPhasePercentage float64, row []string, negative string) {
  for i := range row {
    rowPercentage := float64(i) / float64(len(row))

    if (moonPhasePercentage > rowPercentage) || (moonPhasePercentage < -1 * rowPercentage) {
      row[i] = ":"
    } else {
      row[i] = negative
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

  percentage := (moonPhasePercentage - 4) / 4.0

  fillMoon(percentage, topRow[:], " ")
  fillMoon(percentage, firstMoonRow[:], " ")
  fillMoon(percentage, secondMoonRow[:], " ")
  fillMoon(percentage, thirdMoonRow[:], " ")
  fillMoon(percentage, bottomRow[:], ".")

  b.WriteString("       _..._\n")
  b.WriteString("     ." + strings.Join(topRow[:], "") + ".\n")
  b.WriteString("    :" + strings.Join(firstMoonRow[:], "") + ":\n")
  b.WriteString("    :" + strings.Join(secondMoonRow[:], "") + ":\n")
  b.WriteString("    '" + strings.Join(thirdMoonRow[:], "") + "'\n")
  b.WriteString("      ''" + strings.Join(bottomRow[:], "") + "''\n")

  return b.String()
}
