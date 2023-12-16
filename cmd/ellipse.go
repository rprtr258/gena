package main

import "github.com/rprtr258/gena"

func ellipse() {
	const S = 1024
	dc := gena.NewContext(S, S)
	dc.SetRGBA(0, 0, 0, 0.1)
	for i := 0; i < 360; i += 15 {
		dc.Push()
		dc.RelativeTo(complex(S/2, S/2), func(dc *gena.Context) {
			dc.Rotate(gena.Radians(float64(i)))
		})
		dc.DrawEllipse(complex(S/2, S/2), complex(S*7/16, S/8))
		dc.Fill()
		dc.Pop()
	}
	dc.DrawImageAnchored(gena.LoadImage("cmd/gopher.png"), S/2, S/2, 0.5, 0.5)
	dc.SavePNG("ellipse.png")
}
