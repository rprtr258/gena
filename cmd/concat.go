package main

import (
	"math"

	"github.com/rprtr258/gena"
)

func concat() {
	im1, err := gena.LoadPNG("cmd/baboon.png")
	if err != nil {
		panic(err)
	}

	im2, err := gena.LoadPNG("cmd/gopher.png")
	if err != nil {
		panic(err)
	}

	s1 := im1.Bounds().Size()
	s2 := im2.Bounds().Size()

	width := int(math.Max(float64(s1.X), float64(s2.X)))
	height := s1.Y + s2.Y

	dc := gena.NewContext(width, height)
	dc.DrawImage(im1, 0, 0)
	dc.DrawImage(im2, 0, s1.Y)
	dc.SavePNG("concat.png")
}
