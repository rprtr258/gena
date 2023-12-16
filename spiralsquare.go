package gena

import (
	"image/color"
	"math/rand"
)

// SpiralSquare draws a spiral square images.
func SpiralSquare(
	c Canvas,
	lineWidth float64,
	lineColor color.RGBA,
	squareNum int,
	rectSide, decay float64,
	fg color.RGBA, randColor bool,
) {
	dc := NewContextForRGBA(c.Img())

	sl := rectSide
	theta := rand.Intn(360) + 1
	for i := 0; i < squareNum; i++ {
		dc.Stack(func(ctx *Context) {
			dc.Translate(c.Size() / 2)
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
				dc.SetColor(c.ColorSchema[rand.Intn(len(c.ColorSchema))])
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
