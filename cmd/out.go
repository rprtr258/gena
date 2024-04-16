package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func out() *image.RGBA {
	dc := NewContext(complex(1000, 1000))
	dc.DrawCircle(complex(350, 500), 300)
	dc.Clip()
	dc.DrawCircle(complex(650, 500), 300)
	dc.Clip()
	dc.DrawRectangle(0, complex(1000, 1000))
	dc.SetColor(ColorRGB(0, 0, 0))
	dc.Fill()
	return dc.Image()
}
