package main

import (
	"image"

	"github.com/golang/freetype/truetype"
	. "github.com/rprtr258/gena"
	"golang.org/x/image/font/gofont/goregular"
)

func rotatedText() *image.RGBA {
	const S = 400
	dc := NewContext(complex(S, S))
	dc.SetColor(ColorRGB(1, 1, 1))
	dc.Clear()
	dc.SetColor(ColorRGB(0, 0, 0))
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err.Error())
	}
	dc.SetFontFace(truetype.NewFace(font, &truetype.Options{
		Size: 40,
	}))
	text := "Hello, world!"
	wh := dc.MeasureString(text)
	dc.TransformAdd(Rotate(Radians(10)))
	dc.DrawRectangle(complex(100, 180), wh)
	dc.Stroke()
	dc.DrawStringAnchored(text, complex(100, 180), 0)
	return dc.Image()
}
