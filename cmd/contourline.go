package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// ContourLine  draws a contour line image.
// It uses the perlin noise` to do some flow field.
//   - n: indicates how many lines
func ContourLine(dc *Context, colorSchema []color.RGBA, n int) {
	noise := NewPerlinNoiseDeprecated()

	im := dc.Image()
	dc.SetColor(color.RGBA{0x1a, 0x06, 0x33, 0xFF})
	dc.Clear()
	for range Range(n) {
		cls := RandomItem(colorSchema)
		v := Mul2(RandomV2(), Size(im))

		for range Range(1500) {
			theta := noise.NoiseV2_1(v/800) * PI * 2 * 800
			v += Polar(0.4, theta)

			dc.SetColor(cls)
			dc.DrawEllipse(v, complex(2, 2))
			dc.Fill()

			if v.X() > float64(im.Bounds().Dx()) || v.X() < 0 ||
				v.Y() > float64(im.Bounds().Dy()) || v.Y() < 0 ||
				Random() < 0.001 {
				v = Mul2(RandomV2(), Size(im))
			}
		}
	}
}

func contourline() *image.RGBA {
	dc := NewContext(Diag(1600))
	ContourLine(dc, []color.RGBA{
		{0x58, 0x18, 0x45, 0xFF},
		{0x90, 0x0C, 0x3F, 0xFF},
		{0xC7, 0x00, 0x39, 0xFF},
		{0xFF, 0x57, 0x33, 0xFF},
		{0xFF, 0xC3, 0x0F, 0xFF},
	}, 500)
	return dc.Image()
}
