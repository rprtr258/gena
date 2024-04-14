package main

import (
	"image"
	"math"

	. "github.com/rprtr258/gena"
)

type circles struct {
	pos                V2
	radius             float64
	colorMin, colorMax int
}

// Generative draws a circle with perlin noise.
//   - circleN: The number of circle on this image.
//   - dotsN: The number of dots in each circle.
//   - colorMin: The minimum color.
//   - colorMax: The maximum color.
func PerlinPearls(c *image.RGBA, lineWidth float64, alpha uint8, circleN, dotsN, colorMin, colorMax, iters int) {
	dc := NewContextForRGBA(c)
	dc.SetLineWidth(0.5)
	dc.SetColor(Black)

	cs := make([]circles, 0)
	for len(cs) < circleN {
		c := circles{
			pos: complex(
				RandomFloat64(100, float64(c.Bounds().Dx())-50),
				RandomFloat64(100, float64(c.Bounds().Dy())-50),
			),
			radius:   RandomFloat64(20, 100),
			colorMin: colorMin,
			colorMax: colorMax,
		}
		var overlapping bool
		for _, cl := range cs {
			d := Dist(c.pos, cl.pos)
			if d < c.radius+cl.radius {
				overlapping = true
				break
			}
		}

		if !overlapping {
			cs = append(cs, c)
		}
	}

	ds := make([][]dot, circleN)
	for i := range Range(circleN) {
		dots := make([]dot, dotsN)
		for j := range Range(dotsN) {
			p := cs[i].pos + Polar(cs[i].radius, RandomFloat64(0, math.Pi*2))
			dots[j] = dot{pos: p, prev: p, count: 0}
		}
		ds[i] = dots
	}

	noise := NewPerlinNoiseDeprecated()

	for i := range Range(circleN) {
		dc.SetLineWidth(0.5)
		dc.SetColor(Black)
		dc.DrawCircle(cs[i].pos, cs[i].radius)
		dc.Stroke()

		const factor = 0.008
		for range Range(iters) {
			for k := range ds[i] {
				n := noise.NoiseV2_1(ds[i][k].pos * factor)
				ds[i][k].prev = ds[i][k].pos
				ds[i][k].pos += Polar(2, math.Pi*(n*2+float64(ds[i][k].count)))

				if Dist(cs[i].pos, ds[i][k].pos) > cs[i].radius+1 {
					ds[i][k].count += 1
				}

				if Dist(cs[i].pos, ds[i][k].pos) < cs[i].radius &&
					Dist(cs[i].pos, ds[i][k].prev) < cs[i].radius {
					dc.SetLineWidth(lineWidth)
					rgb := HSV{
						H: int(Remap(n, 0, 1, float64(cs[i].colorMin), float64(cs[i].colorMax))),
						S: 100,
						V: 20,
					}.ToRGB(100, 100, 100)
					rgb.A = alpha
					dc.SetColor(rgb)
					dc.DrawLine(ds[i][k].prev, ds[i][k].pos)
					dc.Stroke()
				}
			}
		}
	}
}
