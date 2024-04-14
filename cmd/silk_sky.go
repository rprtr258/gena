package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

// SilkSky would draw an image with multiple circles converge to one point or one circle.
//   - circleNum: The number of the circles in this drawing.
//   - sunRadius: The radius of the sun. The sun is a point/circle where other circles meet.
func SilkSky(c *image.RGBA, alpha, circleNum int, sunRadius float64) {
	dc := NewContextForRGBA(c)
	m := Mul2(RandomV2(), Size(c))/5 + Size(c)*3/5

	mh := circleNum*2 + 2
	ms := circleNum*2 + 50
	mv := 100

	for i := range Range(circleNum) {
		for j := range Range(circleNum) {
			hsv := HSV{
				H: circleNum + j,
				S: i + 50,
				V: 70,
			}
			rgba := hsv.ToRGB(mh, ms, mv)
			n := Mul2(Plus(complex(float64(i), float64(j)), 0.5), Size(c)) / Coeff(float64(circleNum))
			dc.SetRGBA255(rgba, alpha)
			r := Dist(n, m)
			dc.DrawCircleV2(n, r-sunRadius/2)
			dc.Fill()
		}
	}
}
