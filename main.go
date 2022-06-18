// main package
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// main function.
func main() {
	rx, _ := EfficientToNormalRx("4")
	fmt.Println(rx)
}

type normalRx struct {
	rightSphere          float32
	rightCylinder        float32
	rightAxis            float32
	rightAdd             float32
	rightHorizontalPrism float32
	rightVerticalPrism   float32
	leftSphere           float32
	leftCylinder         float32
	leftAxis             float32
	leftAdd              float32
	leftHorizontalPrism  float32
	leftVerticalPrism    float32
}

// convertDioptre turns the numerical representation of an efficient rx into a
// normal rx.
func convertDioptre(i int64) (f float32, err error) {
	f = float32(i) / 4
	return f, err
}

// "0\n4" -> R plano L +1.00 DS
// "4#" -> R +1.00 DS L +1.00 DS
// "4.-4#110/nx90" -> R +1.00/-1.00x110 L plano/-1.00x90
// "12#.-8#x180#.12# -> R +3.00/-2.00x180 L +3.00/-2.00x180  Add +3.00 
//                                                           Add +3.00

// EfficientToNormalRx turns an efficient rx into a normal Rx struct
func EfficientToNormalRx(eRx string) (nRx normalRx, err error) {
  // split by carriage return for r and l
    r = :strings.split(eRx, "\n")
    if len(r) >
    
	// r := strings.Split(s, ".")
	// if len(r) == 1 {
	// 	i, _ := strconv.ParseInt(r[0], 10, 32)
	//
	// 	n.rightSphere, _ = convertDioptre(i)
	// }
	return n, err
}

// func NormalRxToEfficient() {}
