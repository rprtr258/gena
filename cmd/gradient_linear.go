package main

import (
	"image/color"

	. "github.com/rprtr258/gena"
)

func gradientLinear() {
	dc := NewContext(complex(500, 400))

	grad := PatternGradientLinear(
		complex(20, 320),
		complex(400, 20),
		Stops{0: Green, 0.5: Red, 1: Blue})

	dc.SetColor(color.White)
	dc.DrawRectangle(complex(20, 20), complex(400-20, 300))
	dc.Stroke()

	dc.SetStrokeStyle(grad)
	dc.SetLineWidth(4)
	dc.MoveTo(10 + 10i)
	dc.LineTo(410 + 10i)
	dc.LineTo(410 + 100i)
	dc.LineTo(10 + 100i)
	dc.ClosePath()
	dc.Stroke()

	dc.SetFillStyle(grad)
	dc.MoveTo(10 + 120i)
	dc.LineTo(410 + 120i)
	dc.LineTo(410 + 300i)
	dc.LineTo(10 + 300i)
	dc.ClosePath()
	dc.Fill()

	SavePNG("gradientLinear.png", dc.Image())
}
