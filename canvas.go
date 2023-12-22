package gena

import (
	"image"
	"image/color"
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

type Canvas struct {
	Height, Width int
	img           *image.RGBA
}

// NewCanvas returns a Canvas
func NewCanvas(w, h int) Canvas {
	return Canvas{
		Height: h,
		Width:  w,
		img:    image.NewRGBA(image.Rect(0, 0, w, h)),
	}
}

func (c Canvas) Img() *image.RGBA {
	return c.img
}

func (c Canvas) Size() V2 {
	return complex(float64(c.Width), float64(c.Height))
}

// FillBackground fills the background of the Canvas
func (c Canvas) FillBackground(bg color.RGBA) {
	draw.Draw(c.Img(), c.Img().Bounds(), &image.Uniform{bg}, image.Point{}, draw.Src)
}
