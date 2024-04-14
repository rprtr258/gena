package main

import . "github.com/rprtr258/gena"

func tiling() {
	const NX = 4
	const NY = 3
	im := Load("cmd/gopher.png")
	sz := Size(im)
	dc := NewContext(Mul2(sz, complex(NX, NY)))
	for y := range Range(NY) {
		for x := range Range(NX) {
			pos := complex(float64(x), float64(y))
			dc.DrawImage(im, Mul2(pos, sz))
		}
	}
	SavePNG("tiling.png", dc.Image())
}
