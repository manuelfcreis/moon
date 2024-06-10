package moon

import (
  "testing"
  "time"
  "strings"
)

func removeWhiteSpace(x string) string {
	x = strings.Replace(x, " ", "", -1)
	x = strings.Replace(x, "\t", "", -1)
	x = strings.Replace(x, "\n", "", -1)

  return x
}

func testMoonDrawing(t *testing.T, expectedDrawing string, moonDrawing string) {
  if removeWhiteSpace(expectedDrawing) != removeWhiteSpace(moonDrawing) {
    t.Errorf("Drawing did not generate correctly. \n %s \n %s", expectedDrawing, moonDrawing) 
  }
}

func testMoonInformation(t *testing.T, expectedInformation string, moonInformation string) {
  if expectedInformation != moonInformation {
    t.Errorf("Incorrect moon information. Expected %s but got %s", expectedInformation, moonInformation)
  }

}

func TestFullMoonDrawing(t *testing.T) {
  information, moonDrawing := Moon(time.Date(2024, 05, 23, 0, 0, 0, 0, time.UTC)) 
  expectedDrawing := `
       _..._
     .:::::::.
    :::::::::::
    :::::::::::
    ':::::::::'
      '':::''
  `


  testMoonInformation(t, "Full moon", information)
  testMoonDrawing(t, expectedDrawing, moonDrawing)
}

func TestLastQuarterMoonDrawing(t *testing.T) {
  information, moonDrawing := Moon(time.Date(2024, 6, 28, 0, 0, 0, 0, time.UTC)) 
  expectedDrawing := `
       _..._
     .::::   .
    ::::::    :
    ::::::    :
    ':::::    '
      ''::.''
  `

  testMoonInformation(t, "Last quarter moon", information)
  testMoonDrawing(t, expectedDrawing, moonDrawing)
}

func TestNewMoonDrawing(t *testing.T) {
  information, moonDrawing := Moon(time.Date(2024, 6, 6, 0, 0, 0, 0, time.UTC)) 
  expectedDrawing := `
       _..._
     .       .
    :         :
    :         :
    '.       '
      ''...''
  `


  testMoonInformation(t, "New moon", information)
  testMoonDrawing(t, expectedDrawing, moonDrawing)
}

func TestFirstQuarterMoonDrawing(t *testing.T) {
  information, moonDrawing := Moon(time.Date(2024, 6, 14, 0, 0, 0, 0, time.UTC)) 
  expectedDrawing := `
       _..._
     .    :::.
    :     :::::
    :     :::::
    '.    ::::'
      ''..:''
  `

  testMoonInformation(t, "First quarter moon", information)
  testMoonDrawing(t, expectedDrawing, moonDrawing)
}

