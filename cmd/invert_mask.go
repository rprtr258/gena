package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func invertMask() *image.RGBA {
	dc := NewContext(complex(1024, 1024))
	dc.DrawCircle(complex(512, 512), 384)
	dc.Clip()
	dc.InvertMask()
	dc.DrawRectangle(0, complex(1024, 1024))
	dc.SetColor(ColorRGB(0, 0, 0))
	dc.Fill()
	return dc.Image()
}
