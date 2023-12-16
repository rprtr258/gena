package main

import "github.com/rprtr258/gena"

func patternFill() {
	im, err := gena.LoadPNG("cmd/baboon.png")
	if err != nil {
		panic(err)
	}
	pattern := gena.NewSurfacePattern(im, gena.RepeatBoth)
	dc := gena.NewContext(600, 600)
	dc.MoveTo(20, 20)
	dc.LineTo(590, 20)
	dc.LineTo(590, 590)
	dc.LineTo(20, 590)
	dc.ClosePath()
	dc.SetFillStyle(pattern)
	dc.Fill()
	dc.SavePNG("patternFill.png")
}
