package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// Generative draws a color circle images.
func ColorCircle(im *image.RGBA, colorSchema []color.NRGBA, n int) {
	dc := NewContextFromRGBA(im)

	for range Range(n) {
		v := Mul2(RandomV2N(Diag(-0.1), Diag(1.1)), Size(im))
		s := RandomF64(0, RandomF64(0, X(Size(im))/2)) + 10

		switch RandomWeighted(map[int]float64{0: 4, 1: 4, 2: 1}) {
		case 0:
			// circle made of points
			n := RandomIntN(4, 30)
			cs := RandomF64(2, 8)
			dc.SetColor(RandomItem(colorSchema))
			dc.Stack(func(ctx *Context) {
				for _, a := range RangeF64(0, PI*2, n) {
					dc.DrawCircle(v+Polar(s/2, a), cs/2)
					dc.Fill()
				}
			})
		case 1:
			// regular circle
			dc.SetLineWidth(Random())
			dc.SetColor(RandomItem(colorSchema))
			dc.DrawCircle(v, RandomF64(0, s)/2)
			dc.Stroke()
		case 2:
			// ring cloud of points
			cl := RandomItem(colorSchema)
			dc.SetLineWidth(1.0)
			sx := s * RandomF64(0.1, 0.55)
			for _, j := range RangeStepF64(0.0001, sx, 1) {
				cl.A = 64
				dc.SetColor(cl)

				dd := s + j*2
				for range Range(200) {
					theta := RandomF64(0, PI*2)
					dc.DrawPoint(v+Polar(dd*0.3, theta), 0.6)
					dc.Stroke()
				}
			}
		}
	}
}
