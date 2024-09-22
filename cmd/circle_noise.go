package main

import (
	"image"

	. "github.com/rprtr258/gena"
	"github.com/schollz/progressbar/v3"
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
	dc *Context,
	lineWidth float64,
	aalpha uint8,
	angles []float64,
	colorMin, colorMax float64,
	iters int,
	noise PerlinNoise,
) {
	im := dc.Image()
	dc.SetColor(White)
	dc.Clear()
	dc.SetLineWidth(2.0)
	dc.SetColor(Black)
	radius := float64(im.Bounds().Dx()) * 0.8 / 2
	dc.DrawCircle(Size(im)/2, radius)
	dc.Stroke()
	// ctex.Clip()

	dots := make([]dot, len(angles))
	for i, angle := range angles {
		v := Size(im)/2 + Polar(radius, angle*PI*2)
		dots[i] = dot{pos: v, prev: v, count: 0}
	}

	const factor = 0.008
	for range Range(iters) {
		for i := range dots {
			n := noise.NoiseV2_1(dots[i].pos * factor)
			alpha := PI * (n*2 + float64(dots[i].count))
			nn := Polar(2, alpha)
			dots[i].prev, dots[i].pos = dots[i].pos, dots[i].pos+nn
			if Dist(Size(im)/2, dots[i].pos) > radius+2 {
				dots[i].count += 1
			}

			if Dist(Size(im)/2, dots[i].pos) < radius &&
				Dist(Size(im)/2, dots[i].prev) < radius {
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

func circlenoise() []*image.RGBA {
	const _dots = 2000
	const _iters = 100 // 0
	noise := NewPerlinNoiseDeprecated()
	bar := progressbar.Default(_iters)
	res := make([]*image.RGBA, _iters)
	for i, ralpha := range RangeF64(0, 2*PI, _iters) {
		_ = bar.Add(1)

		angles := make([]float64, _dots)
		for j := range angles {
			angles[j] = noise.NoiseV2_1(P(float64(j), 0)+Polar(1, ralpha)) + ralpha
		}

		peepo := [4096]float64{}
		for j := range peepo {
			peepo[j] = noise.NoiseV2_1(P(float64(j), 1) + Polar(1, ralpha))
		}
		noise2 := NewPerlinNoise(peepo)

		dc := NewContext(Diag(500))
		CircleNoise(dc, 0.3, 80, angles, 60, 80, 400, noise2)
		res[i] = dc.Image()
	}
	return res
}
