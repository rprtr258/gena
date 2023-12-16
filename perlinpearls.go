package gena

import (
	"math"
)

type perlinPearls struct {
	circleN            int
	dotsN              int
	colorMin, colorMax int
	iters              int
}

type circles struct {
	pos                V2
	radius             float64
	colorMin, colorMax int
}

func NewPerlinPerls(circleN, dotsN, colorMin, colorMax, iters int) *perlinPearls {
	return &perlinPearls{
		circleN:  circleN,
		dotsN:    dotsN,
		colorMin: colorMin,
		colorMax: colorMax,
		iters:    iters,
	}
}

// Generative draws a circle with perlin noise.
func (pp *perlinPearls) Generative(c Canvas) {
	ctex := NewContextForRGBA(c.Img())
	ctex.SetLineWidth(0.5)
	ctex.SetColor(Black)

	cs := make([]circles, 0)
	for len(cs) < pp.circleN {
		c := circles{
			pos: complex(
				RandomFloat64(100, float64(c.Width())-50),
				RandomFloat64(100, float64(c.Height())-50),
			),
			radius:   RandomFloat64(20, 100),
			colorMin: pp.colorMin,
			colorMax: pp.colorMax,
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

	ds := make([][]dot, pp.circleN)
	for i := 0; i < pp.circleN; i++ {
		dots := make([]dot, pp.dotsN)
		for j := 0; j < pp.dotsN; j++ {
			p := cs[i].pos + Polar(cs[i].radius, RandomFloat64(0, math.Pi*2))
			dots[j] = dot{pos: p, prev: p, count: 0}
		}
		ds[i] = dots
	}

	noise := NewPerlinNoiseDeprecated()

	for i := 0; i < pp.circleN; i++ {
		ctex.SetLineWidth(0.5)
		ctex.SetColor(Black)
		ctex.DrawCircleV2(cs[i].pos, cs[i].radius)
		ctex.Stroke()

		const factor = 0.008
		for j := 0; j < pp.iters; j++ {
			for k := range ds[i] {
				n := noise.NoiseV2(ds[i][k].pos * factor)
				ds[i][k].prev = ds[i][k].pos
				ds[i][k].pos += Polar(2, math.Pi*(n*2+float64(ds[i][k].count)))

				if Dist(cs[i].pos, ds[i][k].pos) > cs[i].radius+1 {
					ds[i][k].count += 1
				}

				if Dist(cs[i].pos, ds[i][k].pos) < cs[i].radius &&
					Dist(cs[i].pos, ds[i][k].prev) < cs[i].radius {
					ctex.SetLineWidth(c.LineWidth)
					rgb := HSV{
						H: int(Remap(n, 0, 1, float64(cs[i].colorMin), float64(cs[i].colorMax))),
						S: 100,
						V: 20,
					}.ToRGB(100, 100, 100)
					rgb.A = uint8(c.Alpha)
					ctex.SetColor(rgb)
					ctex.DrawLine(ds[i][k].prev, ds[i][k].pos)
					ctex.Stroke()
				}
			}
		}
	}
}
