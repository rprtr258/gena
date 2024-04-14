package main

import . "github.com/rprtr258/gena"

func lines() {
	const W = 1024
	const H = 1024
	dc := NewContext(complex(W, H))
	dc.SetColor(ColorRGB(0, 0, 0))
	dc.Clear()
	for range Range(1000) {
		x1 := Random() * W
		y1 := Random() * H
		x2 := Random() * W
		y2 := Random() * H
		r := Random()
		g := Random()
		b := Random()
		a := Random()*0.5 + 0.5
		w := Random()*4 + 1
		dc.SetColor(ColorRGBA(r, g, b, a))
		dc.SetLineWidth(w)
		dc.DrawLine(complex(x1, y1), complex(x2, y2))
		dc.Stroke()
	}
	SavePNG("lines.png", dc.Image())
}
