package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func mask() *image.RGBA {
	dc := NewContext(complex(512, 512))
	dc.DrawRoundedRectangle(0, complex(512, 512), 64)
	dc.Clip()
	dc.DrawImage(Load("cmd/baboon.png"), 0)
	return dc.Image()
}
