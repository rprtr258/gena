package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func patternFill() *image.RGBA {
	dc := NewContext(complex(600, 600))
	dc.MoveTo(complex(20, 20))
	dc.LineTo(complex(590, 20))
	dc.LineTo(complex(590, 590))
	dc.LineTo(complex(20, 590))
	dc.ClosePath()
	dc.SetFillStyle(PatternSurface(Load("cmd/baboon.png"), RepeatBoth))
	dc.Fill()
	return dc.Image()
}
