package main

import (
	"math"

	. "github.com/rprtr258/gena"
)

func main() {
	const SZ = 1000
	dc := NewContext(Diag(SZ))
	dc.SetColor(Black)
	dc.Clear()
	dc.TransformAdd(
		Translate(Diag(SZ/10)),
		Scale(Diag(SZ/10*8)),
	)
	dc.SetColor(White)
	for range SZ / 2 {
		x := Random()
		const sigma = 8
		y1 := math.Exp(-math.Pow((1-x)*sigma, 1))
		y2 := math.Exp(-math.Pow(x*sigma, 1))
		dc.DrawLine(P(x, x+(1-x)*y2), P(x, x*(1-y1)))
		dc.Stroke()
	}

	SavePNG("test.png", dc.Image())
}
