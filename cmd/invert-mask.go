package main

import . "github.com/rprtr258/gena"

func invertMask() {
	dc := NewContext(complex(1024, 1024))
	dc.DrawCircle(complex(512, 512), 384)
	dc.Clip()
	dc.InvertMask()
	dc.DrawRectangle(0, complex(1024, 1024))
	dc.SetColor(ColorRGB(0, 0, 0))
	dc.Fill()
	SavePNG("invertMask.png", dc.Image())
}
