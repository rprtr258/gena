package main

import "github.com/rprtr258/gena"

func circle() {
	dc := gena.NewContext(complex(1000, 1000))
	dc.DrawCircle(500+500i, 400)
	dc.SetColor(gena.ColorRGB(0, 0, 0))
	dc.Fill()
	gena.SavePNG("circle.png", dc.Image())
}
