package gena

import (
	"math"
	"math/rand"
)

type contourLine struct {
	lineNum int
}

// NewContourLine returns a contourLine object.
func NewContourLine(lineNum int) *contourLine {
	return &contourLine{
		lineNum: lineNum,
	}
}

// Generative draws a contour line image.
func (cl *contourLine) Generative(c Canvas) {
	ctex := NewContextForRGBA(c.Img())
	noise := NewPerlinNoiseDeprecated()
	for i := 0; i < cl.lineNum; i++ {
		cls := c.ColorSchema[rand.Intn(len(c.ColorSchema))]
		v := Mul2(RandomV2(), c.Size())

		for j := 0; j < 1500; j++ {
			theta := noise.NoiseV2(v/800) * math.Pi * 2 * 800
			v += Polar(0.4, theta)

			ctex.SetColor(cls)
			ctex.DrawEllipse(v, complex(2, 2))
			ctex.Fill()

			if X(v) > float64(c.Width()) || X(v) < 0 ||
				Y(v) > float64(c.Height()) || Y(v) < 0 ||
				rand.Float64() < 0.001 {
				v = Mul2(RandomV2(), c.Size())
			}
		}
	}
}
