package main

import "github.com/rprtr258/gena"

func patternFill() {
	dc := gena.NewContext(600, 600)
	dc.MoveTo(complex(20, 20))
	dc.LineTo(complex(590, 20))
	dc.LineTo(complex(590, 590))
	dc.LineTo(complex(20, 590))
	dc.ClosePath()
	dc.SetFillStyle(gena.PatternSurface(gena.LoadPNG("cmd/baboon.png"), gena.RepeatBoth))
	dc.Fill()
	gena.SavePNG("patternFill.png", dc.Image())
}
