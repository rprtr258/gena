package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

var loremIpsumLines = []string{
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
	"tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,",
	"quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo",
	"consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse",
	"cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat",
	"non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
}

func lorem() *image.RGBA {
	const W = 800
	const H = 400
	dc := NewContext(complex(W, H))
	dc.SetColor(ColorRGB(1, 1, 1))
	dc.Clear()
	dc.SetColor(ColorRGB(0, 0, 0))
	// dc.LoadFontFace("/Library/Fonts/Arial.ttf", 18)
	const h = 24
	for i, line := range loremIpsumLines {
		y := H/2 - h*len(loremIpsumLines)/2 + i*h
		dc.DrawStringAnchored(line, complex(400, float64(y)), complex(0.5, 0.5))
	}
	return dc.Image()
}
