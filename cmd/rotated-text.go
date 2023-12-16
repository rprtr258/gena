package main

import (
	"github.com/golang/freetype/truetype"
	"github.com/rprtr258/gena"
	"golang.org/x/image/font/gofont/goregular"
)

func rotatedText() {
	const S = 400
	dc := gena.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic("")
	}
	face := truetype.NewFace(font, &truetype.Options{
		Size: 40,
	})
	dc.SetFontFace(face)
	text := "Hello, world!"
	w, h := dc.MeasureString(text)
	dc.Rotate(gena.Radians(10))
	dc.DrawRectangle(complex(100, 180), complex(w, h))
	dc.Stroke()
	dc.DrawStringAnchored(text, 100, 180, 0.0, 0.0)
	dc.SavePNG("rotatedText.png")
}
