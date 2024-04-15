package main

import (
	"image/color"

	. "github.com/rprtr258/gena"
)

func gradientText() {
	const W = 1024
	const H = 512

	dc := NewContext(complex(W, H))

	// draw text
	dc.SetColor(ColorRGB(0, 0, 0))
	dc.LoadFontFace("/Library/Fonts/Impact.ttf", 128)
	dc.DrawStringAnchored("Gradient Text", complex(W/2, H/2), complex(0.5, 0.5))

	// get the context as an alpha mask
	mask := dc.AsMask()

	// clear the context
	dc.SetColor(ColorRGB(1, 1, 1))
	dc.Clear()

	// set a gradient
	g := PatternGradientLinear(0, complex(W, H), Stops{
		0: color.RGBA{255, 0, 0, 255},
		1: color.RGBA{0, 0, 255, 255},
	})
	dc.SetFillStyle(g)

	// using the mask, fill the context with the gradient
	dc.SetMask(mask)
	dc.DrawRectangle(0, complex(W, H))
	dc.Fill()

	SavePNG("gradientText.png", dc.Image())
}
