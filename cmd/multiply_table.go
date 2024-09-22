package main

import (
	"image"

	. "github.com/rprtr258/gena"
)

func multiplyTable() []*image.RGBA {
	const n = 10
	const FRAMES = 500
	size := Diag(1000)
	const R = 450
	color := White

	imgs := make([]*image.RGBA, FRAMES)
	for i := range Range(FRAMES) {
		dc := NewContext(size)
		dc.SetColor(Black)
		dc.Clear()
		dc.TransformAdd(Translate(size / 2))
		dc.TransformAdd(Scale(Diag(R)))

		k := float64(n*i) / FRAMES
		angle := 2 * PI * k / n
		for i := range Range(n + 1) {
			ang := angle * float64(i)
			dc.LineTo(Polar(1, ang))
		}

		dc.SetColor(color)
		dc.SetLineWidth(2)
		dc.Stroke()
		imgs[i] = dc.Image()
	}

	return imgs
}
