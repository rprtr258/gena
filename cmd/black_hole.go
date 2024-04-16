package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// BlackHole draws a black hole image.
//   - n: number of circles
//   - density: density of circles
//   - circleGap: gap between two circles
func BlackHole(
	dc *Context,
	lineWidth float64,
	lineColor color.RGBA,
	n int,
	density, circleGap float64,
	// random params
	noise PerlinNoise,
	alpha float64, // [0, 1)
) {
	dc.SetColor(color.RGBA{30, 30, 30, 255})
	dc.Clear()

	kMax := Remap(alpha, 0, 1, 0.5, 1)
	dc.SetLineWidth(0.4)
	dc.SetColor(lineColor)

	for i := range Range(n) {
		radius := float64(dc.Image().Bounds().Dx()/10) + float64(i)*0.05
		k := kMax * Sqrt(float64(i)/float64(n))
		noisiness := density * Pow(float64(i)/float64(n), 2)

		base := Size(dc.Image()) / 2
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

func blackhole() *image.RGBA {
	dc := NewContext(Diag(500))
	BlackHole(dc, 1, Tomato, 200, 400, 0.03, NewPerlinNoiseDeprecated(), Random())
	return dc.Image()
}
