package main

import (
	"math/rand"

	. "github.com/rprtr258/gena"
)

func CreatePoints(n int) []V2 {
	points := make([]V2, n)
	for i := range Range(n) {
		x := 0.5 + rand.NormFloat64()*0.1
		y := x + rand.NormFloat64()*0.1
		points[i] = complex(x, y)
	}
	return points
}

func scatter() {
	const S = 1024
	const P = 64
	dc := NewContext(complex(S, S))
	dc.InvertY()
	dc.SetColor(ColorRGB(1, 1, 1))
	dc.Clear()
	points := CreatePoints(1000)
	dc.TransformAdd(Translate(complex(P, P)))
	dc.TransformAdd(Scale(complex(S-P*2, S-P*2)))
	// draw minor grid
	for i := 1; i <= 10; i++ {
		x := float64(i) / 10
		dc.MoveTo(complex(x, 0))
		dc.LineTo(complex(x, 1))
		dc.MoveTo(complex(0, x))
		dc.LineTo(complex(1, x))
	}
	dc.SetColor(ColorRGBA(0, 0, 0, 0.25))
	dc.SetLineWidth(1)
	dc.Stroke()
	// draw axes
	dc.MoveTo(0)
	dc.LineTo(complex(1, 0))
	dc.MoveTo(0)
	dc.LineTo(complex(0, 1))
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
		dc.DrawStringAnchored("Chart Title", complex(S, P)/Coeff(2), complex(0.5, 0.5))
		dc.LoadFontFace("/Library/Fonts/Arial.ttf", 18)
		dc.DrawStringAnchored("X Axis Title", complex(S/2, S-P/2), complex(0.5, 0.5))
	}
	SavePNG("scatter.png", dc.Image())
}
