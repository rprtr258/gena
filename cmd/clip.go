package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func clip() *image.RGBA {
	dc := NewContext(complex(1000, 1000))
	dc.DrawCircle(350+500i, 300)
	dc.Clip()
	dc.DrawCircle(650+500i, 300)
	dc.Clip()
	dc.DrawRectangle(0, 1000+1000i)
	dc.SetColor(ColorRGB(0, 0, 0))
	dc.Fill()
	return dc.Image()
}
