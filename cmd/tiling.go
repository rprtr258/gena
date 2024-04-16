package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func tiling() *image.RGBA {
	const NX = 4
	const NY = 3
	im := Load("cmd/gopher.png")
	sz := Size(im)
	dc := NewContext(Mul2(sz, complex(NX, NY)))
	for _, f := range RangeV2_2(NX, NY) {
		dc.DrawImage(im, Mul2(f, sz))
	}
	return dc.Image()
}
