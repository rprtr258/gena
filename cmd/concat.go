package main

import "github.com/rprtr258/gena"

func concat() {
	im1 := gena.LoadPNG("cmd/baboon.png")
	s1 := im1.Bounds().Size()

	im2 := gena.LoadPNG("cmd/gopher.png")
	s2 := im2.Bounds().Size()

	dc := gena.NewContext(max(s1.X, s2.X), s1.Y+s2.Y)
	dc.DrawImage(im1, 0, 0)
	dc.DrawImage(im2, 0, s1.Y)
	gena.SavePNG("concat.png", dc.Image())
}
