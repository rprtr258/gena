package main

import "github.com/rprtr258/gena"

func concat() {
	im1 := gena.LoadPNG("cmd/baboon.png")
	s1 := gena.Size(im1)

	im2 := gena.LoadPNG("cmd/gopher.png")
	s2 := gena.Size(im2)

	dc := gena.NewContext(complex(max(gena.X(s1), gena.X(s2)), gena.Y(s1)+gena.Y(s2)))
	dc.DrawImage(im1, 0)
	dc.DrawImage(im2, complex(0, gena.Y(s1)))
	gena.SavePNG("concat.png", dc.Image())
}
