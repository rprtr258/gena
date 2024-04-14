package main

import . "github.com/rprtr258/gena"

func unicode() {
	const S = 4096 * 2
	const T = 16 * 2
	const F = 28
	dc := NewContext(complex(S, S))
	dc.SetColor(ColorRGB(1, 1, 1))
	dc.Clear()
	dc.SetColor(ColorRGB(0, 0, 0))
	dc.LoadFontFace("Xolonium-Regular.ttf", F)
	for r := range Range(256) {
		for c := range Range(256) {
			i := r*256 + c
			dc.DrawStringAnchored(string(rune(i)), complex(float64(c), float64(r))*Coeff(T)+Diag(T/2), complex(0.5, 0.5))
		}
	}
	SavePNG("unicode.png", dc.Image())
}
