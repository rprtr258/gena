package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// ContourLine  draws a contour line image.
// It uses the perlin noise` to do some flow field.
//   - n: indicates how many lines
func ContourLine(im *image.RGBA, colorSchema []color.RGBA, n int) {
	dc := NewContextFromRGBA(im)
	noise := NewPerlinNoiseDeprecated()
	for range Range(n) {
		cls := RandomItem(colorSchema)
		v := Mul2(RandomV2(), Size(im))

		for range Range(1500) {
			theta := noise.NoiseV2_1(v/800) * PI * 2 * 800
			v += Polar(0.4, theta)

			dc.SetColor(cls)
			dc.DrawEllipse(v, complex(2, 2))
			dc.Fill()

			if X(v) > float64(im.Bounds().Dx()) || X(v) < 0 ||
				Y(v) > float64(im.Bounds().Dy()) || Y(v) < 0 ||
				Random() < 0.001 {
				v = Mul2(RandomV2(), Size(im))
			}
		}
	}
}
