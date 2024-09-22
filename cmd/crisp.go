package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func crisp() *image.RGBA {
	const W = 1000
	const H = 1000
	const Minor = 10
	const Major = 100

	dc := NewContext(P(W, H))
	dc.SetColor(ColorRGB(1, 1, 1))
	dc.Clear()

	// minor grid
	for x := Minor; x < W; x += Minor {
		fx := float64(x) + 0.5
		dc.DrawLine(P(fx, 0), P(fx, H))
	}
	for y := Minor; y < H; y += Minor {
		fy := float64(y) + 0.5
		dc.DrawLine(P(0, fy), P(W, fy))
	}
	dc.SetLineWidth(1)
	dc.SetColor(ColorRGBA(0, 0, 0, 0.25))
	dc.Stroke()

	// major grid
	for x := Major; x < W; x += Major {
		fx := float64(x) + 0.5
		dc.DrawLine(P(fx, 0), P(fx, H))
	}
	for y := Major; y < H; y += Major {
		fy := float64(y) + 0.5
		dc.DrawLine(P(0, fy), P(W, fy))
	}
	dc.SetLineWidth(1)
	dc.SetColor(ColorRGBA(0, 0, 0, 0.5))
	dc.Stroke()

	// axes
	dc.DrawLine(P(W/2, 0), P(W/2, H))
	dc.DrawLine(P(0, H/2), P(W, H/2))
	dc.SetLineWidth(1)
	dc.SetColor(ColorRGBA(0, 0, 0, 1))
	dc.Stroke()

	return dc.Image()
}
