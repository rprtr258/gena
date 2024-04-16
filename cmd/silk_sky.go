package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

// SilkSky would draw an image with multiple circles converge to one point or one circle.
//   - n: number of circles
//   - sunRadius: radius of sun. The sun is a point/circle where other circles meet.
func SilkSky(dc *Context, alpha, n int, sunRadius float64) {
	im := dc.Image()
	m := Mul2(RandomV2(), Size(im))/5 + Size(im)*3/5

	mh := n*2 + 2
	ms := n*2 + 50
	mv := 100

	for _, f := range RangeV2_2(n, n) {
		hsv := HSV{H: n + int(Y(f)), S: int(X(f)) + 50, V: 70}
		rgba := hsv.ToRGB(mh, ms, mv)
		n := Mul2(f+Diag(0.5), Size(im)) / Coeff(float64(n))
		dc.SetColor(ColorRGBA255(rgba, alpha))
		r := Dist(n, m)
		dc.DrawCircle(n, r-sunRadius/2)
		dc.Fill()
	}
}

func silkSky() *image.RGBA {
	dc := NewContext(Diag(600))
	SilkSky(dc, 10, 15, 5)
	return dc.Image()
}
