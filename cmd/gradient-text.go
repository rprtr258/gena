package main

import (
	"image/color"

	"github.com/rprtr258/gena"
)

const (
	W = 1024
	H = 512
)

func gradientText() {
	dc := gena.NewContext(W, H)

	// draw text
	dc.SetRGB(0, 0, 0)
	dc.LoadFontFace("/Library/Fonts/Impact.ttf", 128)
	dc.DrawStringAnchored("Gradient Text", W/2, H/2, 0.5, 0.5)

	// get the context as an alpha mask
	mask := dc.AsMask()

	// clear the context
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// set a gradient
	g := gena.PatternGradientLinear(0, complex(W, H), gena.Stops{
		0: color.RGBA{255, 0, 0, 255},
		1: color.RGBA{0, 0, 255, 255},
	})
	dc.SetFillStyle(g)

	// using the mask, fill the context with the gradient
	dc.SetMask(mask)
	dc.DrawRectangle(0, complex(W, H))
	dc.Fill()

	gena.SavePNG("gradientText.png", dc.Image())
}
