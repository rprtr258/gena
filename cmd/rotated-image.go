package main

import "github.com/rprtr258/gena"

func rotatedImage() {
	const W = 400
	const H = 500
	im := gena.LoadPNG("cmd/gopher.png")
	iw, ih := im.Bounds().Dx(), im.Bounds().Dy()
	dc := gena.NewContext(W, H)
	// draw outline
	dc.SetColor(gena.ColorHex("#ff0000"))
	dc.SetLineWidth(1)
	dc.DrawRectangle(0, complex(float64(W), float64(H)))
	dc.Stroke()
	// draw full image
	dc.SetColor(gena.ColorHex("#0000ff"))
	dc.SetLineWidth(2)
	dc.DrawRectangle(complex(100, 210), complex(float64(iw), float64(ih)))
	dc.Stroke()
	dc.DrawImage(im, 100, 210)
	// draw image with current matrix applied
	dc.SetColor(gena.ColorHex("#0000ff"))
	dc.SetLineWidth(2)
	dc.Rotate(gena.Radians(10))
	dc.DrawRectangle(complex(100, 0), complex(float64(iw), float64(ih)/2+20.0))
	dc.StrokePreserve()
	dc.Clip()
	dc.DrawImageAnchored(im, 100, 0, 0.0, 0.0)
	gena.SavePNG("rotatedImage.png", dc.Image())
}
