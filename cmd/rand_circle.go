package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"

	. "github.com/rprtr258/gena"
)

type circle1 struct {
	pos    V2
	radius float64
	d      V2
}

func newCircleSlice(cn, w, h int, minStep, maxStep, minRadius, maxRadius float64) []circle1 {
	circles := make([]circle1, 0, cn)
	for range Range(cn) {
		x := rand.Intn(w) + 1
		y := rand.Intn(h) + 1
		radius := float64(rand.Intn(int(minRadius))) + maxRadius - minRadius
		circles = append(circles, circle1{
			pos:    complex(float64(x), float64(y)),
			radius: radius,
			d:      Polar(RandomFloat64(minStep, maxStep), rand.Float64()*math.Pi*2.0),
		})
	}
	return circles
}

func circleSliceUpdate(cs []circle1, bounds image.Rectangle) []circle1 {
	res := make([]circle1, len(cs))
	for i, c := range cs {
		c.pos += c.d

		if X(c.pos) <= 0 {
			c.pos = complex(0, imag(c.pos))
			c.d = complex(-X(c.d), Y(c.d))
		}

		if Y(c.pos) <= 0 {
			c.pos = complex(real(c.pos), 0)
			c.d = complex(X(c.d), -Y(c.d))
		}

		if X(c.pos) > float64(bounds.Dx()) {
			c.pos = complex(float64(bounds.Dx()), imag(c.pos))
			c.d = complex(-X(c.d), Y(c.d))
		}

		if Y(c.pos) > float64(bounds.Dy()) {
			c.pos = complex(real(c.pos), float64(bounds.Dy()))
			c.d = complex(X(c.d), -Y(c.d))
		}

		res[i] = c
	}
	return res
}

// Generative draws a random circles image.
func RandCircle(
	c *image.RGBA,
	colorSchema []color.RGBA,
	lineWidth float64,
	lineColor color.RGBA,
	maxCircle, maxStepsPerCircle int,
	minSteps, maxSteps, minRadius, maxRadius float64,
	isRandColor bool,
	iters int,
) {
	dc := NewContextForRGBA(c)
	for range Range(iters) {
		cn := rand.Intn(maxCircle) + int(maxCircle/3)
		circles := newCircleSlice(cn, c.Bounds().Dx(), c.Bounds().Dy(), minSteps, maxSteps, minRadius, maxRadius)

		for range Range(maxStepsPerCircle) {
			for _, c1 := range circles {
				for _, c2 := range circles {
					cl := lineColor
					if isRandColor {
						cl = colorSchema[rand.Intn(len(colorSchema))]
					}

					if c1 == c2 {
						continue
					}

					distance := Dist(c1.pos, c2.pos)

					if distance <= c1.radius+c2.radius {
						cc := c1.pos + c2.pos/2
						dc.SetColor(ColorRGBA255(cl, 30))
						dc.SetLineWidth(lineWidth)
						dc.DrawCircle(cc, distance/2)
						dc.Stroke()
					}
				}
			}

			circles = circleSliceUpdate(circles, c.Bounds())
		}
	}
}
