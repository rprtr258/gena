package main

import "github.com/rprtr258/gena"

func quadratic() {
	const S = 1000
	dc := gena.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.Translate(complex(S/2, S/2))
	dc.Scale(complex(40, 40))

	const p0 = complex(-10, 0)
	const p1 = complex(-5, -10)
	const p2 = 0
	const p3 = complex(5, 10)
	const p4 = complex(10, 0)

	dc.MoveToV2(p0)
	dc.LineToV2(p1)
	dc.LineToV2(p2)
	dc.LineToV2(p3)
	dc.LineToV2(p4)
	dc.SetHexColor("FF2D00")
	dc.SetLineWidth(8)
	dc.Stroke()

	dc.MoveToV2(p0)
	dc.QuadraticTo(p1, p2)
	dc.QuadraticTo(p3, p4)
	dc.SetHexColor("3E606F")
	dc.SetLineWidth(16)
	dc.FillPreserve()
	dc.SetRGB(0, 0, 0)
	dc.Stroke()

	dc.DrawCircleV2(p0, 0.5)
	dc.DrawCircleV2(p1, 0.5)
	dc.DrawCircleV2(p2, 0.5)
	dc.DrawCircleV2(p3, 0.5)
	dc.DrawCircleV2(p4, 0.5)
	dc.SetRGB(1, 1, 1)
	dc.FillPreserve()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(4)
	dc.Stroke()

	dc.LoadFontFace("/Library/Fonts/Arial.ttf", 200)
	dc.DrawStringAnchored("g", -5, 5, 0.5, 0.5)
	dc.DrawStringAnchored("G", 5, -5, 0.5, 0.5)

	dc.SavePNG("quadratic.png")
}
