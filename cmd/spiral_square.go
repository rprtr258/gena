package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// SpiralSquare draws a spiral square images.
func SpiralSquare(
	im *image.RGBA,
	colorSchema []color.RGBA,
	lineWidth float64,
	lineColor color.RGBA,
	squareNum int,
	rectSide, decay float64,
	fg color.RGBA,
	randColor bool,
) {
	dc := NewContextFromRGBA(im)
	dc.TransformAdd(Translate(Size(im) / 2))

	sl := rectSide
	theta := RandomInt(360) + 1
	for range Range(squareNum) {
		dc.TransformAdd(Rotate(Radians(float64(theta))))
		dc.Stack(func(ctx *Context) {
			dc.TransformAdd(Scale(Diag(sl)))

			dc.LineTo(complex(-0.5, 0.5))
			dc.LineTo(complex(0.5, 0.5))
			dc.LineTo(complex(0.5, -0.5))
			dc.LineTo(complex(-0.5, -0.5))
			dc.LineTo(complex(-0.5, 0.5))

			dc.SetLineWidth(lineWidth)
			dc.SetColor(lineColor)
			dc.StrokePreserve()

			if randColor {
				dc.SetColor(RandomItem(colorSchema))
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
