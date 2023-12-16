package gena

import (
	"math"
	"math/rand"
)

// ContourLine  draws a contour line image.
// It uses the perlin noise` to do some flow field.
//   - lineNum: It indicates how many lines.
func ContourLine(c Canvas, lineNum int) {
	ctex := NewContextForRGBA(c.Img())
	noise := NewPerlinNoiseDeprecated()
	for i := 0; i < lineNum; i++ {
		cls := c.ColorSchema[rand.Intn(len(c.ColorSchema))]
		v := Mul2(RandomV2(), c.Size())

		for j := 0; j < 1500; j++ {
			theta := noise.NoiseV2(v/800) * math.Pi * 2 * 800
			v += Polar(0.4, theta)

			ctex.SetColor(cls)
			ctex.DrawEllipse(v, complex(2, 2))
			ctex.Fill()

			if X(v) > float64(c.Width) || X(v) < 0 ||
				Y(v) > float64(c.Height) || Y(v) < 0 ||
				rand.Float64() < 0.001 {
				v = Mul2(RandomV2(), c.Size())
			}
		}
	}
}
