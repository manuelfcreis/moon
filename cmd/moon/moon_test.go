package moon

import (
  "testing"
)

func TestDrawFullMoon(t *testing.T) {
  moonDrawing := drawMoon(8) 
  expectedDrawing := `
       _..._
     .:::::::.
    :::::::::::
    :::::::::::
    ':::::::::'
      '':::''
  `
  if moonDrawing != expectedDrawing {
    t.Fatalf(`Hello("Gladys") = %q, %v`, moonDrawing, expectedDrawing) 
  }
}
