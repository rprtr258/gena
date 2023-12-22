package main

import "github.com/rprtr258/gena"

func text() {
	const S = 1024
	dc := gena.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.LoadFontFace("/Library/Fonts/Arial.ttf", 96)
	dc.DrawStringAnchored("Hello, world!", S/2, S/2, 0.5, 0.5)
	gena.SavePNG("text.png", dc.Image())
}
