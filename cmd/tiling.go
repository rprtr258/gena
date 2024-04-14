package main

import "github.com/rprtr258/gena"

func tiling() {
	const NX = 4
	const NY = 3
	im := gena.LoadPNG("cmd/gopher.png")
	sz := gena.Size(im)
	dc := gena.NewContext(gena.Mul2(sz, complex(NX, NY)))
	for y := range gena.Range(NY) {
		for x := range gena.Range(NX) {
			pos := complex(float64(x), float64(y))
			dc.DrawImage(im, gena.Mul2(pos, sz))
		}
	}
	gena.SavePNG("tiling.png", dc.Image())
}
