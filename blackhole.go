package gena

import (
	"image"
	"image/color"
	"math"
)

// BlackHole draws a black hole image.
//   - circleN: The number of the circle.
//   - density: Control the density of the circle.
//   - circleGap: Identify the gap between two circles.
func BlackHole(
	c *image.RGBA,
	lineWidth float64,
	lineColor color.RGBA,
	circleN int,
	density, circleGap float64,
	// random params
	noise PerlinNoise,
	alpha float64, // [0, 1)
) {
	dc := NewContextForRGBA(c)
	kMax := Remap(alpha, 0, 1, 0.5, 1)
	dc.SetLineWidth(0.4)
	dc.SetColor(lineColor)

	for i := range circleN {
		radius := float64(c.Bounds().Dx()/10) + float64(i)*0.05
		k := kMax * math.Sqrt(float64(i)/float64(circleN))
		noisiness := density * math.Pow(float64(i)/float64(circleN), 2)

		base := Size(c) / 2
		for theta := 0.0; theta < 2*math.Pi; theta += 2 * math.Pi / 360 {
			r1 := math.Cos(theta) + 1
			r2 := math.Sin(theta) + 1
			r := radius + noise.Noise3D(k*r1, k*r2, float64(i)*circleGap)*noisiness
			dc.LineToV2(base + Polar(r, theta))
		}
		dc.Stroke()
		dc.ClearPath()
	}
}
