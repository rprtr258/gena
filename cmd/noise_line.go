package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// Generative draws a noise line image.
// NoiseLine draws some random line and circles based on `perlin noise`.
//   - n: The number of random line.
func NoiseLine(im *image.RGBA, colorSchema []color.RGBA, iters int) {
	dc := NewContextFromRGBA(im)
	noise := NewPerlinNoiseDeprecated()

	dc.SetColor(Black)
	for range Range(80) {
		x := Random() * float64(im.Bounds().Dx())
		y := Random() * float64(im.Bounds().Dy())

		s := Random() * float64(im.Bounds().Dx()) / 8
		dc.SetLineWidth(0.5)
		dc.DrawEllipse(complex(x, y), complex(s, s))
		dc.Stroke()
	}

	t := Random() * 10
	for range Range(iters) {
		x := RandomF64(-0.5, 1.5) * float64(im.Bounds().Dx())
		y := RandomF64(-0.5, 1.5) * float64(im.Bounds().Dy())
		cl := RandomItem(colorSchema)
		cl.A = 255

		l := 400
		for j := range Range(l) {
			ns := 0.0005
			w := Sin(PI*float64(j)/float64(l-1)) * 5
			theta := noise.Noise3_1(x*ns, y*ns, t) * 100
			dc.SetColor(cl)
			dc.DrawCircle(complex(x, y), w)
			dc.Fill()
			x += Cos(theta)
			y += Sin(theta)
			t += 0.0000003
		}
	}
}
