package main

import . "github.com/rprtr258/gena"

func random() float64 {
	return RandomFloat64(-1, 1)
}

func point() V2 {
	return complex(random(), random())
}

func drawCurve(dc *Context) {
	dc.SetColor(ColorRGBA(0, 0, 0, 0.1))
	dc.FillPreserve()
	dc.SetColor(ColorRGB(0, 0, 0))
	dc.SetLineWidth(12)
	dc.Stroke()
}

func drawPoints(dc *Context) {
	dc.SetColor(ColorRGBA(1, 0, 0, 0.5))
	dc.SetLineWidth(2)
	dc.Stroke()
}

func randomQuadratic(dc *Context) {
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

func randomCubic(dc *Context) {
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
	dc := NewContext(complex(W, H) * Coeff(S))
	dc.SetColor(ColorRGB(1, 1, 1))
	dc.Clear()
	for j := range Range(H) {
		for i := range Range(W) {
			v := S * (complex(float64(i), float64(j)) + Diag(0.5))
			dc.Stack(func(dc *Context) {
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
	SavePNG("beziers.png", dc.Image())
}
