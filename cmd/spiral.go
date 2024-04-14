package main

import (
	"math"

	"github.com/rprtr258/gena"
)

func spiral() {
	const S = 1024
	const N = 2048
	dc := gena.NewContext(complex(S, S))
	dc.SetColor(gena.ColorRGB(1, 1, 1))
	dc.Clear()
	dc.SetColor(gena.ColorRGB(0, 0, 0))
	for i := 0; i <= N; i++ {
		t := float64(i) / N
		d := t*S*0.4 + 10
		a := t * math.Pi * 2 * 20
		dc.DrawCircle(gena.Plus(gena.Polar(d, a), S/2), t*8)
	}
	dc.Fill()
	gena.SavePNG("spiral.png", dc.Image())
}
