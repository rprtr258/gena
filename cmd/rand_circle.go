package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

type circle1 struct {
	center V2
	radius float64
	d      V2
}

func newCircleSlice(cn int, size V2, minStep, maxStep, minRadius, maxRadius float64) []circle1 {
	circles := make([]circle1, cn)
	for i := range circles {
		circles[i] = circle1{
			center: Diag(1) + RandomV2N(0, size),
			radius: RandomF64(minRadius, maxRadius),
			d:      Polar(RandomF64(minStep, maxStep), Random()*PI*2),
		}
	}
	return circles
}

func circleSliceUpdate(cs []circle1, size V2) []circle1 {
	res := make([]circle1, len(cs))
	for i, c := range cs {
		c.center += c.d

		if X(c.center) <= 0 {
			c.center = complex(0, imag(c.center))
			c.d = complex(-X(c.d), Y(c.d))
		}

		if Y(c.center) <= 0 {
			c.center = complex(real(c.center), 0)
			c.d = complex(X(c.d), -Y(c.d))
		}

		if X(c.center) > X(size) {
			c.center = complex(X(size), imag(c.center))
			c.d = complex(-X(c.d), Y(c.d))
		}

		if Y(c.center) > Y(size) {
			c.center = complex(real(c.center), Y(size))
			c.d = complex(X(c.d), -Y(c.d))
		}

		res[i] = c
	}
	return res
}

// Generative draws a random circles image.
func RandCircle(
	dc *Context,
	palette Pattern1D,
	lineWidth float64,
	lineColor color.RGBA,
	maxCircle, maxStepsPerCircle int,
	minSteps, maxSteps, minRadius, maxRadius float64,
	iters int,
) {
	im := dc.Image()
	dc.SetColor(MistyRose)
	dc.Clear()
	for range Range(iters) {
		cn := RandomInt(maxCircle) + int(maxCircle/3)
		circles := newCircleSlice(cn, Size(im), minSteps, maxSteps, minRadius, maxRadius)

		for range Range(maxStepsPerCircle) {
			for _, c1 := range circles {
				for _, c2 := range circles {
					if c1 == c2 {
						continue
					}

					distance := Dist(c1.center, c2.center)
					if distance <= c1.radius+c2.radius {
						cc := c1.center + c2.center/2
						dc.SetColor(palette.Random())
						dc.SetLineWidth(lineWidth)
						dc.DrawCircle(cc, distance/2)
						dc.Stroke()
					}
				}
			}

			circles = circleSliceUpdate(circles, Size(im))
		}
	}
}

func randCircle() *image.RGBA {
	dc := NewContext(Diag(500))
	RandCircle(dc, Plasma, 1, color.RGBA{122, 122, 122, 30}, 30, 80, 0.2, 2, 10, 30, 4)
	return dc.Image()
}
