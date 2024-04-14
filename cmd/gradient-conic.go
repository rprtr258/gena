package main

import (
	"image/color"

	. "github.com/rprtr258/gena"
)

func gradientConic() {
	dc := NewContext(Diag(400))

	grad1 := PatternGradientConic(Diag(200), 0, Stops{
		0.0: color.Black,
		0.5: color.RGBA{255, 215, 0, 255},
		1.0: color.RGBA{255, 0, 0, 255},
	})

	grad2 := PatternGradientConic(Diag(200), 90, Stops{
		0.00: color.RGBA{255, 0, 0, 255},
		0.16: color.RGBA{255, 255, 0, 255},
		0.33: color.RGBA{0, 255, 0, 255},
		0.50: color.RGBA{0, 255, 255, 255},
		0.66: color.RGBA{0, 0, 255, 255},
		0.83: color.RGBA{255, 0, 255, 255},
		1.00: color.RGBA{255, 0, 0, 255},
	})

	dc.SetStrokeStyle(grad1)
	dc.SetLineWidth(20)
	dc.DrawCircle(Diag(200), 180)
	dc.Stroke()

	dc.SetFillStyle(grad2)
	dc.DrawCircle(Diag(200), 150)
	dc.Fill()

	SavePNG("gradient-conic.png", dc.Image())
}
