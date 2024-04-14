package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

// Generative draws a sircle moving images.
//   - n: number of circles
func CircleMove(im *image.RGBA, n int) {
	dc := NewContextFromRGBA(im)
	dc.SetLineWidth(0.3)
	noise := NewPerlinNoiseDeprecated()
	cl := RandomInt(255)
	for i := range Range(n) {
		var cxx float64
		np := 300.0
		for j := 0.0; j < np; j += 1.0 {
			theta := Remap(j, 0, np, 0, PI*2)
			cx := float64(i)*3 - 200.0
			cy := float64(im.Bounds().Dy())/2 + Sin(float64(i)/50)*float64(im.Bounds().Dy())/12.0
			xx := Cos(theta+cx/10) * float64(im.Bounds().Dy()) / 6.0
			yy := Sin(theta+cx/10) * float64(im.Bounds().Dy()) / 6.0
			p := complex(xx, yy)
			xx = (xx + cx) / 150
			yy = (yy + cy) / 150
			p *= Coeff(1 + 1.5*noise.Noise2_1(xx, yy))
			dc.LineTo(p + complex(cx, cy))
			cxx = cx
		}

		hue := int(cxx/4) - cl
		if hue < 0 {
			hue += 255
		}

		rgba := HSV{H: hue, S: 180, V: 120}.ToRGB(255, 255, 255)
		rgba.A = 255
		dc.SetColor(rgba)
		dc.Stroke()
		dc.ClosePath()
	}
}
