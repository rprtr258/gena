package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func CreatePoints(n int) []V2 {
	points := make([]V2, n)
	for i := range points {
		x := RandomGaussian(0.5, 0.1)
		y := RandomGaussian(x, 0.1)
		points[i] = P(x, y)
	}
	return points
}

func scatter() *image.RGBA {
	const S = 1024
	const p = 64
	dc := NewContext(P(S, S))
	dc.InvertY()
	dc.SetColor(ColorRGB(1, 1, 1))
	dc.Clear()
	points := CreatePoints(1000)
	dc.TransformAdd(Translate(P(p, p)))
	dc.TransformAdd(Scale(P(S-p*2, S-p*2)))
	// draw minor grid
	for i := 1; i <= 10; i++ {
		x := float64(i) / 10
		dc.MoveTo(P(x, 0))
		dc.LineTo(P(x, 1))
		dc.MoveTo(P(0, x))
		dc.LineTo(P(1, x))
	}
	dc.SetColor(ColorRGBA(0, 0, 0, 0.25))
	dc.SetLineWidth(1)
	dc.Stroke()
	// draw axes
	dc.MoveTo(0)
	dc.LineTo(P(1, 0))
	dc.MoveTo(0)
	dc.LineTo(P(0, 1))
	dc.SetColor(ColorRGB(0, 0, 0))
	dc.SetLineWidth(4)
	dc.Stroke()
	// draw points
	dc.SetColor(ColorRGBA(0, 0, 1, 0.5))
	for _, p := range points {
		dc.DrawCircle(p, 3.0/S)
		dc.Fill()
	}
	// draw text
	dc.TransformSet(Identity)
	dc.SetColor(ColorRGB(0, 0, 0))
	if false { // TODO: fix font loading
		dc.LoadFontFace("/Library/Fonts/Arial Bold.ttf", 24)
		dc.DrawStringAnchored("Chart Title", P(S, p)/Coeff(2), P(0.5, 0.5))
		dc.LoadFontFace("/Library/Fonts/Arial.ttf", 18)
		dc.DrawStringAnchored("X Axis Title", P(S/2, S-p/2), P(0.5, 0.5))
	}
	return dc.Image()
}
