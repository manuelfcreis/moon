package moon

import (
  "fmt"
  "math"
  "time"
  "strings"
  "github.com/charmbracelet/lipgloss"
)

func calculateMoonPercentage(currentDate time.Time) float64 {
  originalFullMoon := time.Date(2024, 05, 23, 0, 0, 0, 0, time.UTC)
  daysBetweenFullMoons := 29.53
  daysSinceFullMoon := currentDate.Sub(originalFullMoon).Hours() / 24

  return daysSinceFullMoon / daysBetweenFullMoons * 8
}

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

func Moon() {
  moonPhasePercentage := calculateMoonPercentage(time.Now())
  moonPhase := determineMoonPhase(moonPhasePercentage) 
  moonDrawing := drawMoon(moonPhasePercentage)

  yellow := "#ffdd91"
  lightBg := "#b093a3"
  darkBg := "#282e54"

  var style = lipgloss.NewStyle().
      Foreground(lipgloss.AdaptiveColor{Light: "#222222", Dark: "#FAFAFA"}).
      Background(lipgloss.AdaptiveColor{Light: lightBg, Dark: darkBg}).
      BorderStyle(lipgloss.RoundedBorder()).
      BorderForeground(lipgloss.Color(yellow)).
      PaddingTop(2).
      PaddingBottom(2).
      PaddingLeft(4).
      PaddingRight(4)


  fmt.Println(style.Render(lipgloss.JoinHorizontal(lipgloss.Top, moonPhase, moonDrawing)))
}
