package main

import (
	"image"
	"math"

	. "github.com/rprtr258/gena"
)

type dot struct {
	pos   V2
	prev  V2
	count int
}

// CircleNoise draws a circle with perlin noise.
//   - angles: The number of dot of random float64-s
//   - colorMin: minimum color
//   - colorMax: maximum color
func CircleNoise(
	c *image.RGBA,
	lineWidth float64,
	aalpha uint8,
	angles []float64,
	colorMin, colorMax float64,
	iters int,
	noise PerlinNoise,
) {
	dc := NewContextForRGBA(c)
	dc.SetLineWidth(2.0)
	dc.SetColor(Black)
	radius := float64(c.Bounds().Dx()) * 0.8 / 2
	dc.DrawCircle(Size(c)/2, radius)
	dc.Stroke()
	// ctex.Clip()

	dots := make([]dot, len(angles))
	for i, angle := range angles {
		v := Size(c)/2 + Polar(radius, angle*math.Pi*2)
		dots[i] = dot{pos: v, prev: v, count: 0}
	}

	const factor = 0.008
	for range Range(iters) {
		for i := range dots {
			n := noise.NoiseV2_1(dots[i].pos * factor)
			alpha := math.Pi * (n*2 + float64(dots[i].count))
			nn := Polar(2, alpha)
			dots[i].prev, dots[i].pos = dots[i].pos, dots[i].pos+nn
			if Dist(Size(c)/2, dots[i].pos) > radius+2 {
				dots[i].count += 1
			}

			if Dist(Size(c)/2, dots[i].pos) < radius &&
				Dist(Size(c)/2, dots[i].prev) < radius {
				dc.SetLineWidth(lineWidth)
				rgb := HSV{
					H: int(Remap(n, 0, 1, colorMin, colorMax)),
					S: 100,
					V: 20,
				}.ToRGB(100, 100, 100)
				rgb.A = aalpha
				dc.SetColor(rgb)
				dc.DrawLine(dots[i].prev, dots[i].pos)
				dc.Stroke()
			}
		}
	}
}
