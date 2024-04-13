package main

import "github.com/rprtr258/gena"

func tiling() {
	const NX = 4
	const NY = 3
	im := gena.LoadPNG("cmd/gopher.png")
	w := im.Bounds().Size().X
	h := im.Bounds().Size().Y
	dc := gena.NewContext(w*NX, h*NY)
	for y := range gena.Range(NY) {
		for x := range gena.Range(NX) {
			dc.DrawImage(im, x*w, y*h)
		}
	}
	gena.SavePNG("tiling.png", dc.Image())
}
