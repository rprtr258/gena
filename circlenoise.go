package gena

import (
	"math"
)

type dot struct {
	pos   V2
	prev  V2
	count int
}

// CircleNoise draws a circle with perlin noise.
func CircleNoise(c Canvas, dotsN int, colorMin, colorMax float64, iters int) {
	ctex := NewContextForRGBA(c.Img())
	ctex.SetLineWidth(2.0)
	ctex.SetColor(Black)
	radius := float64(c.Width()) * 0.8 / 2
	ctex.DrawCircleV2(c.Size()/2, radius)
	ctex.Stroke()
	// ctex.Clip()

	dots := make([]dot, dotsN)
	for i := 0; i < dotsN; i++ {
		theta := RandomFloat64(0, math.Pi*2)
		v := c.Size()/2 + Polar(radius, theta)
		dots[i] = dot{pos: v, prev: v, count: 0}
	}

	const factor = 0.008
	noise := NewPerlinNoiseDeprecated()
	for j := 0; j < iters; j++ {
		for i := range dots {
			n := noise.NoiseV2(dots[i].pos * factor)
			alpha := math.Pi * (n*2 + float64(dots[i].count))
			nn := Polar(2, alpha)
			dots[i].prev, dots[i].pos = dots[i].pos, dots[i].pos+nn
			if Dist(c.Size()/2, dots[i].pos) > radius+2 {
				dots[i].count += 1
			}

			if Dist(c.Size()/2, dots[i].pos) < radius &&
				Dist(c.Size()/2, dots[i].prev) < radius {
				ctex.SetLineWidth(c.LineWidth)
				rgb := HSV{
					H: int(Remap(n, 0, 1, colorMin, colorMax)),
					S: 100,
					V: 20,
				}.ToRGB(100, 100, 100)
				rgb.A = uint8(c.Alpha)
				ctex.SetColor(rgb)
				ctex.DrawLine(dots[i].prev, dots[i].pos)
				ctex.Stroke()
			}
		}
	}
}
