package main

import "github.com/rprtr258/gena"

func mask() {
	dc := gena.NewContext(complex(512, 512))
	dc.DrawRoundedRectangle(0, complex(512, 512), 64)
	dc.Clip()
	dc.DrawImage(gena.Load("cmd/baboon.png"), 0)
	gena.SavePNG("mask.png", dc.Image())
}
