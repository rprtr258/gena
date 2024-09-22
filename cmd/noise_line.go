package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// Generative draws a noise line image.
// NoiseLine draws some random line and circles based on `perlin noise`.
//   - n: The number of random line.
func NoiseLine(dc *Context, colorSchema []color.RGBA, iters int) {
	dc.SetColor(color.RGBA{0xF0, 0xFE, 0xFF, 0xFF})
	dc.Clear()

	noise := NewPerlinNoiseDeprecated()

	dc.SetColor(Black)
	for range Range(80) {
		f := Mul2(RandomV2(), Size(dc.Image()))

		s := Random() * float64(dc.Image().Bounds().Dx()) / 8
		dc.SetLineWidth(0.5)
		dc.DrawEllipse(f, P(s, s))
		dc.Stroke()
	}

	t := Random() * 10
	for range Range(iters) {
		f := Mul2(RandomV2N(Diag(-0.5), Diag(1.5)), Size(dc.Image()))
		cl := RandomItem(colorSchema)
		cl.A = 255

		l := 400
		for j := range Range(l) {
			ns := 0.0005
			w := Sin(PI*float64(j)/float64(l-1)) * 5
			theta := noise.Noise3_1(f.X()*ns, f.Y()*ns, t) * 100
			dc.SetColor(cl)
			dc.DrawCircle(f, w)
			dc.Fill()
			f += Polar(1, theta)
			t += 0.0000003
		}
	}
}

func noiseline() *image.RGBA {
	dc := NewContext(Diag(1000))
	NoiseLine(dc, []color.RGBA{
		{0x06, 0x7B, 0xC2, 0xFF},
		{0x84, 0xBC, 0xDA, 0xFF},
		{0xEC, 0xC3, 0x0B, 0xFF},
		{0xF3, 0x77, 0x48, 0xFF},
		{0xD5, 0x60, 0x62, 0xFF},
	}, 1000)
	return dc.Image()
}
