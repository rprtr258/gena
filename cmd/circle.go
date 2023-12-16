package main

import "github.com/rprtr258/gena"

func circle() {
	dc := gena.NewContext(1000, 1000)
	dc.DrawCircleV2(500+500i, 400)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SavePNG("circle.png")
}
