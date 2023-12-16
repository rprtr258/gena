package gena

import (
	"math"
)

// BlackHole draws a black hole image.
func BlackHole(
	c Canvas,
	circleN int,
	density, circleGap float64,
	// random params
	noise PerlinNoise,
	alpha float64, // [0, 1)
) {
	ctex := NewContextForRGBA(c.Img())
	kMax := Remap(alpha, 0, 1, 0.5, 1)
	ctex.SetLineWidth(0.4)
	ctex.SetColor(c.LineColor)

	for i := 0; i < circleN; i++ {
		radius := float64(c.Width()/10) + float64(i)*0.05
		k := kMax * math.Sqrt(float64(i)/float64(circleN))
		noisiness := density * math.Pow(float64(i)/float64(circleN), 2)

		base := c.Size() / 2
		for theta := 0.0; theta < 2*math.Pi; theta += 2 * math.Pi / 360 {
			r1 := math.Cos(theta) + 1
			r2 := math.Sin(theta) + 1
			r := radius + noise.Noise3D(k*r1, k*r2, float64(i)*circleGap)*noisiness
			ctex.LineToV2(base + Polar(r, theta))
		}
		ctex.Stroke()
		ctex.ClearPath()
	}
}
