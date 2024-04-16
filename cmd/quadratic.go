package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func quadratic() *image.RGBA {
	const S = 1000
	dc := NewContext(complex(S, S))
	dc.SetColor(ColorRGB(1, 1, 1))
	dc.Clear()
	dc.TransformAdd(Translate(complex(S/2, S/2)))
	dc.TransformAdd(Scale(complex(40, 40)))

	const p0 = complex(-10, 0)
	const p1 = complex(-5, -10)
	const p2 = 0
	const p3 = complex(5, 10)
	const p4 = complex(10, 0)

	dc.MoveTo(p0)
	dc.LineTo(p1)
	dc.LineTo(p2)
	dc.LineTo(p3)
	dc.LineTo(p4)
	dc.SetColor(ColorHex("FF2D00"))
	dc.SetLineWidth(8)
	dc.Stroke()

	dc.MoveTo(p0)
	dc.QuadraticTo(p1, p2)
	dc.QuadraticTo(p3, p4)
	dc.SetColor(ColorHex("3E606F"))
	dc.SetLineWidth(16)
	dc.FillPreserve()
	dc.SetColor(ColorRGB(0, 0, 0))
	dc.Stroke()

	dc.DrawCircle(p0, 0.5)
	dc.DrawCircle(p1, 0.5)
	dc.DrawCircle(p2, 0.5)
	dc.DrawCircle(p3, 0.5)
	dc.DrawCircle(p4, 0.5)
	dc.SetColor(ColorRGB(1, 1, 1))
	dc.FillPreserve()
	dc.SetColor(ColorRGB(0, 0, 0))
	dc.SetLineWidth(4)
	dc.Stroke()

	dc.LoadFontFace("/Library/Fonts/Arial.ttf", 200)
	dc.DrawStringAnchored("g", complex(-5, 5), complex(0.5, 0.5))
	dc.DrawStringAnchored("G", complex(5, -5), complex(0.5, 0.5))

	return dc.Image()
}
