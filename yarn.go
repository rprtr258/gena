package gena

import "image/color"

// Yarn draws a yarn image.
//   - n: The number of the curve.
func Yarn(c Canvas, lineWidth float64, lineColor color.RGBA, n int) {
	ctex := NewContextForRGBA(c.Img())
	ctex.SetLineWidth(lineWidth)
	ctex.SetColor(lineColor)
	noise := NewPerlinNoiseDeprecated()

	offset := 0.0
	inc := 0.005
	for i := 0; i < n; i++ {
		p0 := Mul2(complex(
			noise.Noise1D(offset+15),
			noise.Noise1D(offset+55),
		), c.Size())
		p1 := Mul2(complex(
			noise.Noise1D(offset+25),
			noise.Noise1D(offset+65),
		), c.Size())
		p2 := Mul2(complex(
			noise.Noise1D(offset+35),
			noise.Noise1D(offset+75),
		), c.Size())
		p3 := Mul2(complex(
			noise.Noise1D(offset+45),
			noise.Noise1D(offset+85),
		), c.Size())
		ctex.MoveToV2(p0)
		ctex.CubicTo(p1, p2, p3)
		ctex.Stroke()
		ctex.ClearPath()
		offset += inc
	}
}
