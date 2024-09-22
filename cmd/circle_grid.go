package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// CircleGrid draws a circle grid image
func CircleGrid(dc *Context, colorSchema []color.RGBA, lineWidth float64, n int) {
	dc.SetColor(color.RGBA{0xDF, 0xEB, 0xF5, 0xFF})
	dc.Clear()

	W, H := Size(dc.Image()).XY()

	// grid
	{
		const cells = 100

		dc.SetColor(White)
		dc.SetLineWidth(0.6)
		for _, z := range RangeF64(0, W, cells) {
			dc.DrawLine(P(0, z), P(W, z))
			dc.Stroke()
			dc.DrawLine(P(z, 0), P(z, H))
			dc.Stroke()
		}
	}

	w := W / float64(n)

	for _, f := range RangeV2_2(n, n) {
		center := f*Coeff(w) + Diag(w/2)
		dc.Stack(func(dc *Context) {
			dc.TransformAdd(Translate(center))
			dc.SetColor(RandomItem(colorSchema))
			dc.DrawCircle(0, w/2*RandomF64(0.1, 0.5))
			dc.Fill()

			dc.TransformAdd(Rotate(float64(RandomInt(10))))
			dc.SetColor(RandomItem(colorSchema))
			dc.SetLineWidth(lineWidth)

			r := w / 2 * RandomF64(0.6, 0.95)
			switch RandomInt(4) {
			case 0:
				// just orbit
				dc.DrawCircle(0, r)
				dc.Stroke()
			case 1:
				// orbit
				dc.DrawCircle(0, r)
				dc.Stroke()

				// satellites
				n := RandomIntN(1, 4) * 2
				for range Range(n) {
					dc.TransformAdd(Rotate(PI * 2 / float64(n)))
					dc.DrawCircle(P(r, 0), r*0.1)
					dc.Fill()
				}
			case 2:
				// loading-like circles
				n := RandomIntN(8, 20)
				// total arc is quarter of circle up to 4 times
				arc := PI / 2 * float64(RandomIntN(1, 5))
				for _, d := range RangeF64(0, 1, n) {
					dc.TransformAdd(Rotate(arc / float64(n)))
					dc.DrawCircle(P(r/2, 0), min(d, r*0.1)*2)
					dc.Fill()
				}
			case 3:
				// rays
				n := RandomIntN(5, 20)
				theta := PI * 2 / float64(n)
				for range Range(n) {
					dc.TransformAdd(Rotate(theta))
					dc.DrawLine(P(r*0.5, 0), P(r*0.63, 0))
					dc.Stroke()
				}
			}
		})
	}
}

func circleGrid() *image.RGBA {
	dc := NewContext(Diag(500))
	CircleGrid(dc, []color.RGBA{
		{0xED, 0x34, 0x41, 0xFF},
		{0xFF, 0xD6, 0x30, 0xFF},
		{0x32, 0x9F, 0xE3, 0xFF},
		{0x15, 0x42, 0x96, 0xFF},
		{0x00, 0x00, 0x00, 0xFF},
		{0xFF, 0xFF, 0xFF, 0xFF},
	}, 2, RandomIntN(4, 6))
	return dc.Image()
}
