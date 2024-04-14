package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"

	. "github.com/rprtr258/gena"
)

// ContourLine  draws a contour line image.
// It uses the perlin noise` to do some flow field.
//   - lineNum: It indicates how many lines.
func ContourLine(c *image.RGBA, colorSchema []color.RGBA, lineNum int) {
	dc := NewContextForRGBA(c)
	noise := NewPerlinNoiseDeprecated()
	for range Range(lineNum) {
		cls := colorSchema[rand.Intn(len(colorSchema))]
		v := Mul2(RandomV2(), Size(c))

		for range Range(1500) {
			theta := noise.NoiseV2_1(v/800) * math.Pi * 2 * 800
			v += Polar(0.4, theta)

			dc.SetColor(cls)
			dc.DrawEllipse(v, complex(2, 2))
			dc.Fill()

			if X(v) > float64(c.Bounds().Dx()) || X(v) < 0 ||
				Y(v) > float64(c.Bounds().Dy()) || Y(v) < 0 ||
				rand.Float64() < 0.001 {
				v = Mul2(RandomV2(), Size(c))
			}
		}
	}
}
