package main

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/rprtr258/gena"
)

// SpiralSquare draws a spiral square images.
func SpiralSquare(
	c *image.RGBA, colorSchema []color.RGBA,
	lineWidth float64,
	lineColor color.RGBA,
	squareNum int,
	rectSide, decay float64,
	fg color.RGBA, randColor bool,
) {
	dc := gena.NewContextForRGBA(c)

	sl := rectSide
	theta := rand.Intn(360) + 1
	for i := range gena.Range(squareNum) {
		dc.Stack(func(ctx *gena.Context) {
			dc.Translate(gena.Size(c) / 2)
			dc.Rotate(gena.Radians(float64(theta * (i + 1))))

			dc.Scale(complex(sl, sl))

			dc.LineTo(complex(-0.5, 0.5))
			dc.LineTo(complex(0.5, 0.5))
			dc.LineTo(complex(0.5, -0.5))
			dc.LineTo(complex(-0.5, -0.5))
			dc.LineTo(complex(-0.5, 0.5))

			dc.SetLineWidth(lineWidth)
			dc.SetColor(lineColor)
			dc.StrokePreserve()

			if randColor {
				dc.SetColor(colorSchema[rand.Intn(len(colorSchema))])
			} else {
				dc.SetColor(fg)
			}
			dc.Fill()
		})
		sl -= decay * rectSide

		if sl < 0 {
			return
		}
	}
}
