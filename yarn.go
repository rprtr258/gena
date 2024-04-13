package gena

import (
	"image"
	"image/color"
)

// Yarn draws a yarn image.
//   - n: The number of the curve.
func Yarn(c *image.RGBA, lineWidth float64, lineColor color.RGBA, n int) {
	dc := NewContextForRGBA(c)
	dc.SetLineWidth(lineWidth)
	dc.SetColor(lineColor)
	noise := NewPerlinNoiseDeprecated()

	offset := 0.0
	inc := 0.005
	for range Range(n) {
		dc.MoveToV2(Mul2(noise.NoiseV2D2(offset+15, offset+55), Size(c)))
		dc.CubicTo(
			Mul2(noise.NoiseV2D2(offset+25, offset+65), Size(c)),
			Mul2(noise.NoiseV2D2(offset+35, offset+75), Size(c)),
			Mul2(noise.NoiseV2D2(offset+45, offset+85), Size(c)),
		)
		dc.Stroke()
		dc.ClearPath()
		offset += inc
	}
}
