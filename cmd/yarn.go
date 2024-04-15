package main

import (
	"image"
	"image/color"

	. "github.com/rprtr258/gena"
)

// Yarn draws a yarn image.
//   - n: The number of the curve.
func Yarn(im *image.RGBA, lineWidth float64, lineColor color.RGBA, n int) {
	dc := NewContextFromRGBA(im)
	dc.SetLineWidth(lineWidth)
	dc.SetColor(lineColor)
	noise := NewPerlinNoiseDeprecated()

	offset := 0.0
	inc := 0.005
	for range Range(n) {
		dc.MoveTo(Mul2(noise.Noise2_V2(offset+15, offset+55), Size(im)))
		dc.CubicTo(
			Mul2(noise.Noise2_V2(offset+25, offset+65), Size(im)),
			Mul2(noise.Noise2_V2(offset+35, offset+75), Size(im)),
			Mul2(noise.Noise2_V2(offset+45, offset+85), Size(im)),
		)
		dc.Stroke()
		dc.ClearPath()
		offset += inc
	}
}
