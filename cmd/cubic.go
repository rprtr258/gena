package main

import "github.com/rprtr258/gena"

func cubic() {
	const S = 1000
	dc := gena.NewContext(S, S)
	dc.SetColor(gena.ColorRGB(1, 1, 1))
	dc.Clear()
	dc.Translate(complex(S/2, S/2))
	dc.Scale(complex(40, 40))

	p0 := complex(-10, 0)
	p1 := complex(-8, -8)
	p2 := complex(8, 8)
	p3 := complex(10, 0)

	dc.MoveTo(p0)
	dc.CubicTo(p1, p2, p3)
	dc.SetColor(gena.ColorRGBA(0, 0, 0, 0.2))
	dc.SetLineWidth(8)
	dc.FillPreserve()
	dc.SetColor(gena.ColorRGB(0, 0, 0))
	dc.SetDash(16, 24)
	dc.Stroke()

	dc.MoveTo(p0)
	dc.LineTo(p1)
	dc.LineTo(p2)
	dc.LineTo(p3)
	dc.SetColor(gena.ColorRGBA(1, 0, 0, 0.4))
	dc.SetLineWidth(2)
	dc.SetDash(4, 8, 1, 8)
	dc.Stroke()

	gena.SavePNG("cubic.png", dc.Image())
}
