package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// BlackHole draws a black hole image.
//   - circleN: The number of the circle.
//   - density: Control the density of the circle.
//   - circleGap: Identify the gap between two circles.
func BlackHole(
	im *image.RGBA,
	lineWidth float64,
	lineColor color.RGBA,
	circleN int,
	density, circleGap float64,
	// random params
	noise PerlinNoise,
	alpha float64, // [0, 1)
) {
	dc := NewContextFromRGBA(im)
	kMax := Remap(alpha, 0, 1, 0.5, 1)
	dc.SetLineWidth(0.4)
	dc.SetColor(lineColor)

	for i := range Range(circleN) {
		radius := float64(im.Bounds().Dx()/10) + float64(i)*0.05
		k := kMax * Sqrt(float64(i)/float64(circleN))
		noisiness := density * Pow(float64(i)/float64(circleN), 2)

		base := Size(im) / 2
		for theta := 0.0; theta < 2*PI; theta += 2 * PI / 360 {
			r1 := Cos(theta) + 1
			r2 := Sin(theta) + 1
			r := radius + noise.Noise3_1(k*r1, k*r2, float64(i)*circleGap)*noisiness
			dc.LineTo(base + Polar(r, theta))
		}
		dc.Stroke()
		dc.ClearPath()
	}
}
