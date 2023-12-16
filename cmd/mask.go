package main

import "github.com/rprtr258/gena"

func mask() {
	dc := gena.NewContext(512, 512)
	dc.DrawRoundedRectangle(0, complex(512, 512), 64)
	dc.Clip()
	dc.DrawImage(gena.LoadImage("cmd/baboon.png"), 0, 0)
	dc.SavePNG("mask.png")
}
