package main_test

import (
	"testing"

	. "efficientrx"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEfficientrx(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Efficientrx Suite")
}

var _ = Describe("Efficientrx", func() {

	Context("Error Handling", func() {

		It(">1 carriage returns", func() {
			input := "\n\n"
			Expect(func() {
				_, err := EfficientToNormalRx(input)
				Expect(err).To(HaveOccurred())
			}).To(
				PanicWith(
					MatchRegexp(`Maximum carriage return allowed is one.`)))
		})

	})

	Context("Test cases", func() {

		It("R plano\nL +1.00DS", func() {
			input := "0\n4"
			output, err := EfficientToNormalRx(input)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(output).To(Equal(
				NormalRx{
					RightSphere:          0,
					RightCylinder:        0,
					RightAxis:            0,
					RightAdd:             0,
					RightExtraAdds:       nil,
					RightHorizontalPrism: 0,
					RightVerticalPrism:   0,
					LeftSphere:           1.0,
					LeftCylinder:         0,
					LeftAxis:             0,
					LeftAdd:              0,
					LeftExtraAdds:        nil,
					LeftHorizontalPrism:  0,
					LeftVerticalPrism:    0,
				}))
		})

		It("R +1.00\nL +1.00DS", func() {
			input := "4#"
			output, err := EfficientToNormalRx(input)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(output).To(Equal(
				NormalRx{
					RightSphere:          1.0,
					RightCylinder:        0,
					RightAxis:            0,
					RightAdd:             0,
					RightExtraAdds:       nil,
					RightHorizontalPrism: 0,
					RightVerticalPrism:   0,
					LeftSphere:           1.0,
					LeftCylinder:         0,
					LeftAxis:             0,
					LeftAdd:              0,
					LeftExtraAdds:        nil,
					LeftHorizontalPrism:  0,
					LeftVerticalPrism:    0,
				}))
		})

	})
})

// TODO:
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
