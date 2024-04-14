package main

import (
	"image"
	"math"
	"math/rand"

	. "github.com/rprtr258/gena"
)

// Generative draws a sircle moving images.
//   - circleNum: The number of the circle in the image.
func CircleMove(c *image.RGBA, circleNum int) {
	dc := NewContextFromRGBA(c)
	dc.SetLineWidth(0.3)
	noise := NewPerlinNoiseDeprecated()
	cl := rand.Intn(255)
	for i := range Range(circleNum) {
		var cxx float64
		np := 300.0
		for j := 0.0; j < np; j += 1.0 {
			theta := Remap(j, 0, np, 0, math.Pi*2)
			cx := float64(i)*3 - 200.0
			cy := float64(c.Bounds().Dy())/2 + math.Sin(float64(i)/50)*float64(c.Bounds().Dy())/12.0
			xx := math.Cos(theta+cx/10) * float64(c.Bounds().Dy()) / 6.0
			yy := math.Sin(theta+cx/10) * float64(c.Bounds().Dy()) / 6.0
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
