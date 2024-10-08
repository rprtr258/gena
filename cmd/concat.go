package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func concat() *image.RGBA {
	im1 := Load("cmd/baboon.png")
	s1 := Size(im1)

	im2 := Load("cmd/gopher.png")
	s2 := Size(im2)

	dc := NewContext(P(max(s1.X(), s2.X()), s1.Y()+s2.Y()))
	dc.DrawImage(im1, 0)
	dc.DrawImage(im2, P(0, s1.Y()))
	return dc.Image()
}
