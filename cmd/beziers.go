package main

import "github.com/rprtr258/gena"

func random() float64 {
	return gena.RandomFloat64(-1, 1)
}

func point() gena.V2 {
	return complex(random(), random())
}

func drawCurve(dc *gena.Context) {
	dc.SetColor(gena.ColorRGBA(0, 0, 0, 0.1))
	dc.FillPreserve()
	dc.SetColor(gena.ColorRGB(0, 0, 0))
	dc.SetLineWidth(12)
	dc.Stroke()
}

func drawPoints(dc *gena.Context) {
	dc.SetColor(gena.ColorRGBA(1, 0, 0, 0.5))
	dc.SetLineWidth(2)
	dc.Stroke()
}

func randomQuadratic(dc *gena.Context) {
	p0 := point()
	p1 := point()
	p2 := point()
	dc.MoveTo(p0)
	dc.QuadraticTo(p1, p2)
	drawCurve(dc)
	dc.MoveTo(p0)
	dc.LineTo(p1)
	dc.LineTo(p2)
	drawPoints(dc)
}

func randomCubic(dc *gena.Context) {
	p0 := point()
	p1 := point()
	p2 := point()
	p3 := point()
	dc.MoveTo(p0)
	dc.CubicTo(p1, p2, p3)
	drawCurve(dc)
	dc.MoveTo(p0)
	dc.LineTo(p1)
	dc.LineTo(p2)
	dc.LineTo(p3)
	drawPoints(dc)
}

func beziers() {
	const S = 256
	const W = 8
	const H = 8
	dc := gena.NewContext(complex(W, H) * gena.Coeff(S))
	dc.SetColor(gena.ColorRGB(1, 1, 1))
	dc.Clear()
	for j := range gena.Range(H) {
		for i := range gena.Range(W) {
			v := S * (complex(float64(i), float64(j)) + gena.Diag(0.5))
			dc.Stack(func(dc *gena.Context) {
				dc.Translate(v)
				dc.Scale(complex(S/2, S/2))
				if j%2 == 0 {
					randomCubic(dc)
				} else {
					randomQuadratic(dc)
				}
			})
		}
	}
	gena.SavePNG("beziers.png", dc.Image())
}
