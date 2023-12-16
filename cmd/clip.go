package main

import "github.com/rprtr258/gena"

func clip() {
	dc := gena.NewContext(1000, 1000)
	dc.DrawCircleV2(350+500i, 300)
	dc.Clip()
	dc.DrawCircleV2(650+500i, 300)
	dc.Clip()
	dc.DrawRectangle(0, 1000+1000i)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SavePNG("clip.png")
}
