package main

import (
	"image/color"

	"github.com/rprtr258/gena"
)

func gradientLinear() {
	dc := gena.NewContext(500, 400)

	grad := gena.PatternGradientLinear(
		complex(20, 320),
		complex(400, 20),
		gena.Stops{
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
	dc.LineToV2(410 + 10i)
	dc.LineToV2(410 + 100i)
	dc.LineToV2(10 + 100i)
	dc.ClosePath()
	dc.Stroke()

	dc.SetFillStyle(grad)
	dc.MoveTo(10 + 120i)
	dc.LineToV2(410 + 120i)
	dc.LineToV2(410 + 300i)
	dc.LineToV2(10 + 300i)
	dc.ClosePath()
	dc.Fill()

	gena.SavePNG("gradientLinear.png", dc.Image())
}
