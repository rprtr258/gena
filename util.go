package gena

import (
	"image"
	"image/draw"
)

func must(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func must1[T any](res T, err error) T {
	if err != nil {
		panic(err.Error())
	}

	return res
}

func imageToRGBA(src image.Image) *image.RGBA {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)
	draw.Draw(dst, bounds, src, bounds.Min, draw.Src)
	return dst
}

func Range(n int) []struct{} {
	return make([]struct{}, n)
}
