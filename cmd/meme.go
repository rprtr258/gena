package main

import "github.com/rprtr258/gena"

func meme() {
	const S = 1024
	dc := gena.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.LoadFontFace("/Library/Fonts/Impact.ttf", 96)
	dc.SetRGB(0, 0, 0)
	s := "ONE DOES NOT SIMPLY"
	n := 6 // "stroke" size
	for dy := -n; dy <= n; dy++ {
		for dx := -n; dx <= n; dx++ {
			if dx*dx+dy*dy >= n*n {
				// give it rounded corners
				continue
			}
			x := S/2 + float64(dx)
			y := S/2 + float64(dy)
			dc.DrawStringAnchored(s, x, y, 0.5, 0.5)
		}
	}
	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(s, S/2, S/2, 0.5, 0.5)
	gena.SavePNG("meme.png", dc.Image())
}
