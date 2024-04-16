package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func circle() *image.RGBA {
	const SZ = 1000
	dc := NewContext(Diag(SZ))
	dc.DrawCircle(Diag(SZ)/2, SZ*0.4)
	dc.SetColor(Black)
	dc.Fill()
	return dc.Image()
}
