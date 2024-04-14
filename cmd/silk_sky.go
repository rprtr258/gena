package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

// SilkSky would draw an image with multiple circles converge to one point or one circle.
//   - n: number of circles
//   - sunRadius: radius of sun. The sun is a point/circle where other circles meet.
func SilkSky(im *image.RGBA, alpha, n int, sunRadius float64) {
	dc := NewContextFromRGBA(im)
	m := Mul2(RandomV2(), Size(im))/5 + Size(im)*3/5

	mh := n*2 + 2
	ms := n*2 + 50
	mv := 100

	for i := range Range(n) {
		for j := range Range(n) {
			hsv := HSV{H: n + j, S: i + 50, V: 70}
			rgba := hsv.ToRGB(mh, ms, mv)
			n := Mul2(complex(float64(i), float64(j))+Diag(0.5), Size(im)) / Coeff(float64(n))
			dc.SetColor(ColorRGBA255(rgba, alpha))
			r := Dist(n, m)
			dc.DrawCircle(n, r-sunRadius/2)
			dc.Fill()
		}
	}
}
