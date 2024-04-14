package gena

import "image"

type BlendMode int

const (
	Add BlendMode = iota
)

func Blend(src, dest *image.RGBA, mode BlendMode) *image.RGBA {
	img := image.NewRGBA(src.Bounds())
	for i := range Range(src.Bounds().Max.X) {
		for j := range Range(src.Bounds().Max.Y) {
			switch mode {
			case Add:
				if src.RGBAAt(i, j) == Black {
					img.Set(i, j, dest.RGBAAt(i, j))
				} else {
					img.SetRGBA(i, j, Mix(src.RGBAAt(i, j), dest.RGBAAt(i, j)))
				}
			}
		}
	}
	return img
}
