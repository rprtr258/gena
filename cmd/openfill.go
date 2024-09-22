package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func openfill() *image.RGBA {
	const N = 10
	const SZ = 1000
	dc := NewContext(Diag(SZ))
	dc.TransformAdd(Translate(Diag(SZ / N / 2)))
	for _, y := range RangeF64(0, SZ, N) {
		for _, x := range RangeF64(0, SZ, N) {
			f := P(x, y)
			a1 := PI * Random() * 2
			a2 := PI * (Random() + 0.5)
			dc.DrawArc(f, 40, a1, a1+a2)
		}
	}
	dc.SetColor(Black)
	dc.FillPreserve()

	dc.SetColor(White)
	dc.SetLineWidth(8)
	dc.StrokePreserve()

	dc.SetColor(Red)
	dc.SetLineWidth(4)
	dc.StrokePreserve()

	return dc.Image()
}
