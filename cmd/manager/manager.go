package manager

import (
  "flag"
  "fmt"
  "time"
  "moon/cmd/moon"
  "github.com/charmbracelet/lipgloss"
)

func style() lipgloss.Style {
  yellow := "#ffdd91"
  lightBg := "#b093a3"
  darkBg := "#282e54"

  return lipgloss.NewStyle().
      Foreground(lipgloss.AdaptiveColor{Light: "#222222", Dark: "#FAFAFA"}).
      Background(lipgloss.AdaptiveColor{Light: lightBg, Dark: darkBg}).
      BorderStyle(lipgloss.RoundedBorder()).
      BorderForeground(lipgloss.Color(yellow)).
      PaddingTop(2).
      PaddingBottom(2).
      PaddingLeft(4).
      PaddingRight(4)
}

func prepareMoon(targetTime time.Time) string {
  style := style() 
  moonPhase, moonDrawing := moon.Moon(targetTime)
  styledMoon := style.Render(lipgloss.JoinHorizontal(lipgloss.Top,
                                                     moonPhase,
                                                     moonDrawing))

  return styledMoon
}

func Run() {
  var targetTime time.Time
  var nextFullMoon = flag.Bool("next", false, "Use this flag to search for the next full moon")

  flag.Parse()

  if *nextFullMoon {
    targetTime = moon.DaysToFullMoon()
  } else {
    targetTime = time.Now()
  }

  fmt.Println(prepareMoon(targetTime))
}
