package gena

import (
	"image"
	"image/color"
	"math/rand"
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
	dc := NewContextForRGBA(c)

	sl := rectSide
	theta := rand.Intn(360) + 1
	for i := range squareNum {
		dc.Stack(func(ctx *Context) {
			dc.Translate(Size(c) / 2)
			dc.Rotate(Radians(float64(theta * (i + 1))))

			dc.Scale(complex(sl, sl))

			dc.LineTo(-0.5, 0.5)
			dc.LineTo(0.5, 0.5)
			dc.LineTo(0.5, -0.5)
			dc.LineTo(-0.5, -0.5)
			dc.LineTo(-0.5, 0.5)

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
