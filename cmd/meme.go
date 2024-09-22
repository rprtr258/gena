package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func meme() *image.RGBA {
	const S = 1024
	dc := NewContext(P(S, S))
	dc.SetColor(ColorRGB(1, 1, 1))
	dc.Clear()
	dc.LoadFontFace("/Library/Fonts/Impact.ttf", 96)
	dc.SetColor(ColorRGB(0, 0, 0))
	s := "ONE DOES NOT SIMPLY"
	n := 6 // "stroke" size
	for dy := -n; dy <= n; dy++ {
		for dx := -n; dx <= n; dx++ {
			if dx*dx+dy*dy >= n*n {
				// give it rounded corners
				continue
			}
			dc.DrawStringAnchored(s, P(float64(dx), float64(dy))+Diag(S/2), Diag(0.5))
		}
	}
	dc.SetColor(ColorRGB(1, 1, 1))
	dc.DrawStringAnchored(s, P(S/2, S/2), P(0.5, 0.5))
	return dc.Image()
}
