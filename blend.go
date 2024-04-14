package gena

import (
	"image"
	"image/color"
)

type BlendMode int

const (
	Add BlendMode = iota
)

func Blend(src, dest image.Image, mode BlendMode) image.Image {
	img := image.NewRGBA(src.Bounds())
	for i := range Range(src.Bounds().Max.X) {
		for j := range Range(src.Bounds().Max.Y) {
			switch mode {
			case Add:
				if r, g, b, a := src.At(i, j).RGBA(); (color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}) == Black {
					img.Set(i, j, dest.At(i, j))
				} else {
					img.Set(i, j, Mix(src.At(i, j), dest.At(i, j)))
				}
			}
		}
	}
	return img
}
