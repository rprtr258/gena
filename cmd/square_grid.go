package main

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/rprtr258/gena"
)

// Generative draws a grid squares image.
func GirdSquares(c *image.RGBA, colorSchema []color.RGBA, step, rectSize int, decay float64, iters int) {
	dc := gena.NewContextForRGBA(c)

	for x := 0; x < c.Bounds().Dx(); x += step {
		for y := 0; y < c.Bounds().Dy(); y += step {
			cl := colorSchema[rand.Intn(len(colorSchema))]

			v0 := complex(float64(x), float64(y))
			s := float64(rectSize)

			theta := rand.Intn(360) + 1
			for i := range gena.Range(iters) {
				dc.Stack(func(ctx *gena.Context) {
					dc.Translate(gena.Plus(v0, float64(step)/2))
					dc.Rotate(gena.Radians(float64(theta * i)))

					dc.Scale(complex(s, s))

					dc.LineTo(-0.5, 0.5)
					dc.LineTo(0.5, 0.5)
					dc.LineTo(0.5, -0.5)
					dc.LineTo(-0.5, -0.5)
					dc.LineTo(-0.5, 0.5)

					dc.SetLineWidth(3)
					dc.SetColor(gena.Tomato)
					dc.StrokePreserve()
					dc.SetColor(gena.ColorRGBA255(cl, 255))
					dc.Fill()
				})
				s -= decay * float64(rectSize)
			}
		}
	}
}
