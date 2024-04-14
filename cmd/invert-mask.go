package main

import "github.com/rprtr258/gena"

func invertMask() {
	dc := gena.NewContext(1024, 1024)
	dc.DrawCircle(512, 512, 384)
	dc.Clip()
	dc.InvertMask()
	dc.DrawRectangle(0, complex(1024, 1024))
	dc.SetColor(gena.ColorRGB(0, 0, 0))
	dc.Fill()
	gena.SavePNG("invertMask.png", dc.Image())
}
