// main package
package main

import (
	"fmt"
	"strings"
    "strconv"
)

// main function.
func main() {
  rx, _ := EfficientToNormalRx("4")
  fmt.Println(rx)
}

type normalRx struct {
  rightSphere float32
  rightCylinder float32
  rightAxis float32
  rightAdd float32
  rightHorizontalPrism float32
  rightVerticalPrism float32
  leftSphere float32
  leftCylinder float32
  leftAxis float32
  leftAdd float32
  leftHorizontalPrism float32
  leftVerticalPrism float32
}

// convertDioptre turns the numerical representation of an efficient rx into a
// normal rx.
func convertDioptre(i int64) (f float32, err error) {
  f = float32(i) / 4
  return f, err
}

// turn 12.-3x70.8i8d -> "+3.00/-2.00 x 7 8 BaseIn 8 BaseDown"

// EfficientToNormalRx turns an efficient rx into a normal Rx struct
func EfficientToNormalRx(s string) (n normalRx, err error) {
    r := strings.Split(s, ".")
    if len(r) == 1 {
    i, _ := strconv.ParseInt(r[0], 10, 32)

      n.rightSphere, _ = convertDioptre(i)
  }
	return n, err
}

// func NormalRxToEfficient() {}
