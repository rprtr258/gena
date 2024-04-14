package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"

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
		x := rand.Float64() * float64(im.Bounds().Dx())
		y := rand.Float64() * float64(im.Bounds().Dy())

		s := rand.Float64() * float64(im.Bounds().Dx()) / 8
		dc.SetLineWidth(0.5)
		dc.DrawEllipse(complex(x, y), complex(s, s))
		dc.Stroke()
	}

	t := rand.Float64() * 10
	for range Range(iters) {
		x := RandomF64(-0.5, 1.5) * float64(im.Bounds().Dx())
		y := RandomF64(-0.5, 1.5) * float64(im.Bounds().Dy())
		cl := colorSchema[rand.Intn(len(colorSchema))]
		cl.A = 255

		l := 400
		for j := range Range(l) {
			ns := 0.0005
			w := math.Sin(math.Pi*float64(j)/float64(l-1)) * 5
			theta := noise.Noise3_1(x*ns, y*ns, t) * 100
			dc.SetColor(cl)
			dc.DrawCircle(complex(x, y), w)
			dc.Fill()
			x += math.Cos(theta)
			y += math.Sin(theta)
			t += 0.0000003
		}
	}
}
