package main

import (
	"fmt"
	"strconv"
	"strings"
)

// main function.
func main() {
	rx, _ := EfficientToNormalRx("0\n4")
	fmt.Println(rx)
}

type NormalRx struct {
	RightSphere          float32
	RightCylinder        float32
	RightAxis            float32
	RightAdd             float32
	RightExtraAdds       map[float32]int
	RightHorizontalPrism float32
	RightVerticalPrism   float32
	LeftSphere           float32
	LeftCylinder         float32
	LeftAxis             float32
	LeftAdd              float32
	LeftExtraAdds        map[float32]int
	LeftHorizontalPrism  float32
	LeftVerticalPrism    float32
}

// convertDioptre turns the numerical representation of an efficient rx into a
// normal rx.
func convertDioptre(s *string) (f float32, err error) {
	var i int64
	i, err = strconv.ParseInt(*s, 10, 64)
	f = float32(i) / 4
	return f, err
}

// "0\n4" -> R plano
//           L +1.00 DS
// "4#" -> R +1.00 DS
//         L +1.00 DS
// "4.-4#110\nx90" -> R +1.00/-1.00x110
//                    L plano/-1.00x90
// "12#.-8#x180#.12# -> R +3.00/-2.00x180  Add: +3.00
//                      L +3.00/-2.00x180  Add: +3.00
// "27.-5x35.6\n30.-11x85.8" -> R +6.75/-1.25x35 Add: +1.50
//                              L +7.50/-2.75x85 Add: +2.00
// "4.8i12u\n30.12.-3x70.8i.8d" -> R +1.00 DS 2 In 3 Up
//                                 L +3.00/-0.75x70 2 In 2 Down
// "0#.6#" -> R plano Add: +1.50
//            L plano Add: +1.50

// EfficientToNormalRx turns an efficient rx into a normal Rx struct
func EfficientToNormalRx(eRx string) (nRx NormalRx, err error) {
	var (
		rRx  string
		lRx  string
		rSph string
		lSph string
	)
	// split by carriage return for R and L eyes
	// and check there is one or no carriage returns
	rl := strings.Split(eRx, "\n")
	if len(rl) > 2 {
		panic("Maximum carriage return allowed is one.")
	}
	rRx = rl[0]
	if len(rl) == 2 {
		lRx = rl[1]
	}
	// Split Right Eye by `.`
	rElems := strings.Split(rRx, ".")
	rSph = rElems[0]
	// Check if hash #
	lastChar := string(rSph[len(rSph)-1])
	if lastChar == "#" {
		rSph = strings.Replace(rSph, "#", "", 1)
		lSph = rSph
	}
	nRx.RightSphere, err = convertDioptre(&rSph)
	if err != nil {
		_ = fmt.Errorf("Error: %v", err)
	}
	if lRx != "" {
		lElems := strings.Split(lRx, ".")
		if lSph == "" {
			lSph = lElems[0]
			lastChar := string(lSph[len(lSph)-1])
			if lastChar == "#" {
				panic("Left eye cannot contain '#'.")
			}
		}
	}
	nRx.LeftSphere, err = convertDioptre(&lSph)
	if err != nil {
		_ = fmt.Errorf("Error: %v", err)
	}
	return nRx, err
}
