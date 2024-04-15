package main

import . "github.com/rprtr258/gena"

func circle() {
	const SZ = 1000
	dc := NewContext(Diag(SZ))
	dc.DrawCircle(Diag(SZ)/2, SZ*0.4)
	dc.SetColor(Black)
	dc.Fill()
	SavePNG("circle.png", dc.Image())
}
