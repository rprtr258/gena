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
		Stops{
			0:   color.RGBA{0, 255, 0, 255},
			0.5: color.RGBA{255, 0, 0, 255},
			1:   color.RGBA{0, 0, 255, 255},
		})

	dc.SetColor(color.White)
	dc.DrawRectangle(20+20i, 400-20+300i)
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
