package main

import . "github.com/rprtr258/gena"

func ellipse() {
	const S = 1024
	dc := NewContext(Diag(S))
	dc.SetColor(ColorRGBA(0, 0, 0, 0.1))
	dc.TransformAdd(Translate(Diag(S) / 2))
	for _, a := range RangeF64(0, 2*PI, 24) {
		dc.Stack(func(dc *Context) {
			dc.TransformAdd(Rotate(a))
			dc.DrawEllipse(0, complex(S*7/16, S/8))
			dc.Fill()
		})
	}
	dc.DrawImageAnchored(Load("cmd/gopher.png"), 0, complex(0.5, 0.5))
	SavePNG("ellipse.png", dc.Image())
}
