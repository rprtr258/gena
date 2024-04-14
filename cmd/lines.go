package main

import (
	"math/rand"

	"github.com/rprtr258/gena"
)

func lines() {
	const W = 1024
	const H = 1024
	dc := gena.NewContext(complex(W, H))
	dc.SetColor(gena.ColorRGB(0, 0, 0))
	dc.Clear()
	for range gena.Range(1000) {
		x1 := rand.Float64() * W
		y1 := rand.Float64() * H
		x2 := rand.Float64() * W
		y2 := rand.Float64() * H
		r := rand.Float64()
		g := rand.Float64()
		b := rand.Float64()
		a := rand.Float64()*0.5 + 0.5
		w := rand.Float64()*4 + 1
		dc.SetColor(gena.ColorRGBA(r, g, b, a))
		dc.SetLineWidth(w)
		dc.DrawLine(complex(x1, y1), complex(x2, y2))
		dc.Stroke()
	}
	gena.SavePNG("lines.png", dc.Image())
}
