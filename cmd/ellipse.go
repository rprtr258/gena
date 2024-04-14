package main

import . "github.com/rprtr258/gena"

func ellipse() {
	const S = 1024
	dc := NewContext(complex(S, S))
	dc.SetColor(ColorRGBA(0, 0, 0, 0.1))
	for i := 0; i < 360; i += 15 {
		dc.Stack(func(dc *Context) {
			dc.RelativeTo(complex(S/2, S/2), func(dc *Context) {
				dc.TransformAdd(Rotate(Radians(float64(i))))
			})
			dc.DrawEllipse(complex(S/2, S/2), complex(S*7/16, S/8))
			dc.Fill()
		})
	}
	dc.DrawImageAnchored(Load("cmd/gopher.png"), complex(S, S)/Coeff(2), complex(0.5, 0.5))
	SavePNG("ellipse.png", dc.Image())
}
