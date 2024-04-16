package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func text() *image.RGBA {
	const S = 1024
	dc := NewContext(complex(S, S))
	dc.SetColor(ColorRGB(1, 1, 1))
	dc.Clear()
	dc.SetColor(ColorRGB(0, 0, 0))
	dc.LoadFontFace("/Library/Fonts/Arial.ttf", 96)
	dc.DrawStringAnchored("Hello, world!", complex(S, S)/Coeff(2), complex(0.5, 0.5))
	return dc.Image()
}
