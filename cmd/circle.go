package main

import . "github.com/rprtr258/gena"

func circle() {
	dc := NewContext(complex(1000, 1000))
	dc.DrawCircle(500+500i, 400)
	dc.SetColor(ColorRGB(0, 0, 0))
	dc.Fill()
	SavePNG("circle.png", dc.Image())
}
