package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func cubic() *image.RGBA {
	const S = 1000
	dc := NewContext(P(S, S))
	dc.SetColor(ColorRGB(1, 1, 1))
	dc.Clear()
	dc.TransformAdd(Translate(P(S/2, S/2)))
	dc.TransformAdd(Scale(P(40, 40)))

	p0 := P(-10, 0)
	p1 := P(-8, -8)
	p2 := P(8, 8)
	p3 := P(10, 0)

	dc.MoveTo(p0)
	dc.CubicTo(p1, p2, p3)
	dc.SetColor(ColorRGBA(0, 0, 0, 0.2))
	dc.SetLineWidth(8)
	dc.FillPreserve()
	dc.SetColor(ColorRGB(0, 0, 0))
	dc.SetDash(16, 24)
	dc.Stroke()

	dc.MoveTo(p0)
	dc.LineTo(p1)
	dc.LineTo(p2)
	dc.LineTo(p3)
	dc.SetColor(ColorRGBA(1, 0, 0, 0.4))
	dc.SetLineWidth(2)
	dc.SetDash(4, 8, 1, 8)
	dc.Stroke()

	return dc.Image()
}
