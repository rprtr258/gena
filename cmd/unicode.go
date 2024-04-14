package main

import "github.com/rprtr258/gena"

func unicode() {
	const S = 4096 * 2
	const T = 16 * 2
	const F = 28
	dc := gena.NewContext(complex(S, S))
	dc.SetColor(gena.ColorRGB(1, 1, 1))
	dc.Clear()
	dc.SetColor(gena.ColorRGB(0, 0, 0))
	dc.LoadFontFace("Xolonium-Regular.ttf", F)
	for r := range gena.Range(256) {
		for c := range gena.Range(256) {
			i := r*256 + c
			dc.DrawStringAnchored(string(rune(i)), gena.Plus(complex(float64(c), float64(r))*gena.Coeff(T), T/2), complex(0.5, 0.5))
		}
	}
	gena.SavePNG("unicode.png", dc.Image())
}
