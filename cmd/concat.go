package main

import . "github.com/rprtr258/gena"

func concat() {
	im1 := LoadPNG("cmd/baboon.png")
	s1 := Size(im1)

	im2 := LoadPNG("cmd/gopher.png")
	s2 := Size(im2)

	dc := NewContext(complex(max(X(s1), X(s2)), Y(s1)+Y(s2)))
	dc.DrawImage(im1, 0)
	dc.DrawImage(im2, complex(0, Y(s1)))
	SavePNG("concat.png", dc.Image())
}
