package main

import (
	"image/color"

	"github.com/rprtr258/gena"
)

func gradientRadial() {
	dc := gena.NewContext(complex(400, 200))

	grad := gena.PatternGradientRadial(100+100i, 10, 100+120i, 80, gena.Stops{
		0: color.RGBA{0, 255, 0, 255},
		1: color.RGBA{0, 0, 255, 255},
	})

	dc.SetFillStyle(grad)
	dc.DrawRectangle(0, complex(200, 200))
	dc.Fill()

	dc.SetColor(color.White)
	dc.DrawCircle(complex(100, 100), 10)
	dc.Stroke()
	dc.DrawCircle(complex(100, 120), 80)
	dc.Stroke()

	gena.SavePNG("gradientRadial.png", dc.Image())
}
