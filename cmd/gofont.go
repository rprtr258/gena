package main

import (
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/rprtr258/gena"
	"golang.org/x/image/font/gofont/goregular"
)

func gofont() {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	face := truetype.NewFace(font, &truetype.Options{Size: 48})

	dc := gena.NewContext(complex(1024, 1024))
	dc.SetFontFace(face)
	dc.SetColor(gena.ColorRGB(1, 1, 1))
	dc.Clear()
	dc.SetColor(gena.ColorRGB(0, 0, 0))
	dc.DrawStringAnchored("Hello, world!", complex(512, 512), complex(0.5, 0.5))
	gena.SavePNG("gofont.png", dc.Image())
}
