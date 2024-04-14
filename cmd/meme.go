package main

import "github.com/rprtr258/gena"

func meme() {
	const S = 1024
	dc := gena.NewContext(complex(S, S))
	dc.SetColor(gena.ColorRGB(1, 1, 1))
	dc.Clear()
	dc.LoadFontFace("/Library/Fonts/Impact.ttf", 96)
	dc.SetColor(gena.ColorRGB(0, 0, 0))
	s := "ONE DOES NOT SIMPLY"
	n := 6 // "stroke" size
	for dy := -n; dy <= n; dy++ {
		for dx := -n; dx <= n; dx++ {
			if dx*dx+dy*dy >= n*n {
				// give it rounded corners
				continue
			}
			dc.DrawStringAnchored(s, gena.Plus(complex(float64(dx), float64(dy)), S/2), complex(0.5, 0.5))
		}
	}
	dc.SetColor(gena.ColorRGB(1, 1, 1))
	dc.DrawStringAnchored(s, complex(S/2, S/2), complex(0.5, 0.5))
	gena.SavePNG("meme.png", dc.Image())
}
