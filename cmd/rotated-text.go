package main

import (
	"github.com/golang/freetype/truetype"
	"github.com/rprtr258/gena"
	"golang.org/x/image/font/gofont/goregular"
)

func rotatedText() {
	const S = 400
	dc := gena.NewContext(S, S)
	dc.SetColor(gena.ColorRGB(1, 1, 1))
	dc.Clear()
	dc.SetColor(gena.ColorRGB(0, 0, 0))
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err.Error())
	}
	dc.SetFontFace(truetype.NewFace(font, &truetype.Options{
		Size: 40,
	}))
	text := "Hello, world!"
	wh := dc.MeasureString(text)
	dc.Rotate(gena.Radians(10))
	dc.DrawRectangle(complex(100, 180), wh)
	dc.Stroke()
	dc.DrawStringAnchored(text, 100, 180, 0.0, 0.0)
	gena.SavePNG("rotatedText.png", dc.Image())
}
